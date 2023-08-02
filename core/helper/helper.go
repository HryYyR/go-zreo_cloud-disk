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
	rand "math/rand"
	"net/http"
	"net/smtp"
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
func GenerateToken(id uint64, identity, name string) (string, error) {
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 定义过期时间 1day
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
func DecryptToekn(token string) (*define.UserClaim, error) {
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

// 上传文件到oss
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
