package helper

import (
	"cloud_disk/go-zreo_cloud-disk/core/define"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/md5"
	random "crypto/rand"
	"crypto/tls"
	"errors"
	"fmt"
	"math"
	rand "math/rand"
	"mime/multipart"
	"net/http"
	"net/smtp"
	"os"
	"path"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
)

// md5加密
func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func generateEcdsaPrivateKey() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(elliptic.P256(), random.Reader)
}

// token
func GenerateToken(id uint64, identity, name string, expiretime time.Duration) (string, error) {
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiretime)), // 定义过期时间 单位:分钟
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	ttoken, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		fmt.Println("jwt error: ", err)
		return "", err
	}
	return ttoken, nil
}

// 解密token
func DecryptToken(token string) (*define.UserClaim, error) {
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(tk *jwt.Token) (any, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("token is invalid")
	}
	return uc, nil
}

// 发送验证码邮件
func MailSendCode(mail string, code string) error {
	e := email.NewEmail()
	e.From = "zero-cloud-disk <2452719312@qq.com>"
	e.To = []string{mail}
	e.Subject = "Code"                              //标题
	e.HTML = []byte("<h1>你的验证码: " + code + "</h1>") //内容
	err := e.SendWithTLS("smtp.qq.com:465", smtp.PlainAuth("", "2452719312@qq.com", define.EmailPassword, "smtp.qq.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
	if err != nil {
		return err
	}
	return nil
}

// 生成随机验证码
func RandCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return code
}

// 生成随机UUID
func GenerateUUID() string {
	u1 := uuid.NewV4()
	return u1.String()
}

// 直接上传文件到oss
func OssUpload(r *http.Request) (string, error) {
	// 连接oss
	client, err := oss.New(define.OssEndPoint, define.OssaccessKeyID, define.OssaccessSecret)
	if err != nil {
		return "", err
	}

	// 选择桶
	bucket, err := client.Bucket("zero-cloud-disk")
	if err != nil {
		return "", err
	}

	file, fileheader, err := r.FormFile("file") //获取文件信息
	if err != nil {
		return "", err
	}

	objectKey := "cloud-disk/" + GenerateUUID() + path.Ext(fileheader.Filename) //文件名

	err = bucket.PutObject(objectKey, file)
	if err != nil {
		return "", err
	}
	filepath := define.OssPath + "/" + objectKey
	return filepath, nil
}

// 分片上传件到oss
func OssPartUpload(r *http.Request, b []byte, fileheader *multipart.FileHeader) (string, error) {
	// 步骤0：初始化oss
	client, err := oss.New(define.OssEndPoint, define.OssaccessKeyID, define.OssaccessSecret)
	if err != nil {
		return "", err
	}
	bucket, err := client.Bucket("zero-cloud-disk")
	if err != nil {
		return "", nil
	}
	// end
	Ext := path.Ext(fileheader.Filename)
	UUID := GenerateUUID()
	key := "cloud-disk/" + UUID + Ext   //存放云端的路径
	localname := "../img/" + UUID + Ext //存放本地的路径

	myFile, err := os.OpenFile(localname, os.O_CREATE|os.O_RDWR, os.ModePerm) //保存到本地文件
	myFile.Write(b)
	defer myFile.Close()
	defer os.Remove(localname)
	// 步骤1：初始化一个分片上传事件。
	v, err := bucket.InitiateMultipartUpload(key)
	if err != nil {
		return "", err
	}

	chunkNum := math.Ceil(float64(fileheader.Size) / float64(define.ChunkSize)) //被分成几片
	chunks, err := oss.SplitFileByPartNum(localname, int(chunkNum))             //分片文件数组
	if err != nil {
		return "", err
	}
	var parts []oss.UploadPart //用于保存成功上传的分片

	// 步骤2：上传分片
	// todo 改为并发上传
	// var wg sync.WaitGroup
	// wg.Add(int(chunkNum))
	for _, chunk := range chunks {
		// go func(chunk oss.FileChunk, i int) {
		myFile.Seek(chunk.Offset, 0)
		part, err := bucket.UploadPart(v, myFile, chunk.Size, chunk.Number)
		if err != nil {
			return "", err
		}
		parts = append(parts, part)
		// wg.Done()
		// }(chunk, index)
	}
	// wg.Wait()

	// 指定Object的读写权限为私有，默认为继承Bucket的读写权限。
	// objectAcl := oss.ObjectACL(oss.ACLPrivate)

	// 步骤3：完成分片上传。
	// cmur, err := bucket.CompleteMultipartUpload(v, parts, objectAcl)
	cmur, err := bucket.CompleteMultipartUpload(v, parts)
	if err != nil {
		return "", err
	}
	fmt.Println("上传结果:", cmur)
	filepath := define.OssPath + "/" + key
	return filepath, nil
}
