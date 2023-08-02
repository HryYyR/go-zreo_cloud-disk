package test

import (
	"bytes"
	"cloud_disk/go-zreo_cloud-disk/core/define"
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
