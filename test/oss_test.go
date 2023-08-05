package test

import (
	"bytes"
	"cloud_disk/go-zreo_cloud-disk/core/define"
	"fmt"
	"math"
	"os"
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// 通过url上传
func TestFileUpLoadByFilePATH(t *testing.T) {
	client, err := oss.New("oss-cn-chengdu.aliyuncs.com", define.OssaccessKeyID, define.OssaccessSecret)
	if err != nil {
		t.Fatal(err)
	}
	bucket, err := client.Bucket("zero-cloud-disk")
	if err != nil {
		t.Fatal(err)
	}
	err = bucket.PutObjectFromFile("cloud-disk/test.jpeg", "../img/bilbil.jpeg")
	if err != nil {
		t.Fatal(err)
	}
}

// 通过读写流上传
func TestFileUpLoadByReader(t *testing.T) {
	client, err := oss.New("oss-cn-chengdu.aliyuncs.com", define.OssaccessKeyID, define.OssaccessSecret)
	if err != nil {
		t.Fatal(err)
	}
	bucket, err := client.Bucket("zero-cloud-disk")
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.ReadFile("../img/bilbil.jpeg")
	if err != nil {
		t.Fatal(err)
	}

	err = bucket.PutObject("cloud-disk/666.jpeg", bytes.NewReader(f))
	if err != nil {
		t.Fatal(err)
	}
}

// 分片上传初始化 和 上传
func TestInitPartUpload(t *testing.T) {
	// 步骤0：初始化oss
	client, err := oss.New("oss-cn-chengdu.aliyuncs.com", define.OssaccessKeyID, define.OssaccessSecret)
	if err != nil {
		t.Fatal(err)
	}
	bucket, err := client.Bucket("zero-cloud-disk")
	if err != nil {
		t.Fatal(err)
	}
	// end

	key := "cloud-disk/testprotupload.jpg" //存放路径

	// 步骤1：初始化一个分片上传事件。
	v, err := bucket.InitiateMultipartUpload(key)
	if err != nil {
		t.Fatal(err)
	}
	UploadId := v.UploadID //122748E1F5DC4C81B857489BC6C81C32
	fmt.Println(UploadId)

	fd, err := os.Open("../img/test.jpg")                          //源文件
	fi, err := os.Stat("../img/test.jpg")                          //源文件的属性
	chunkNum := math.Ceil(float64(fi.Size()) / float64(chunkSize)) //被分成几片

	chunks, err := oss.SplitFileByPartNum("../img/test.jpg", int(chunkNum)) //分片文件数组
	if err != nil {
		t.Fatal(err)
	}
	var parts []oss.UploadPart //用于保存成功上传的分片

	// 步骤2：上传分片
	// todo 改为并发上传
	// var wg sync.WaitGroup
	// wg.Add(int(chunkNum))
	for _, chunk := range chunks {
		// go func(chunk oss.FileChunk, i int) {
		fd.Seek(chunk.Offset, 0)
		part, err := bucket.UploadPart(v, fd, chunk.Size, chunk.Number)
		if err != nil {
			return
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
		t.Fatal(err)
	}
	fmt.Println("上传结果:", cmur)
}
