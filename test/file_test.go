package test

import (
	"crypto/md5"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"testing"
)

// 分片大小
const chunkSize = 1 * 1024 * 1024 //1mb

// 文件分片
func TestGenerateChunkFile(t *testing.T) {
	fileinfo, err := os.Stat("../img/test.jpg")
	if err != nil {
		t.Fatal(err)
	}
	chunkNum := math.Ceil(float64(fileinfo.Size()) / float64(chunkSize)) //分片个数
	myFile, err := os.OpenFile("../img/test.jpg", os.O_RDONLY, 0666)     //读取文件
	defer myFile.Close()
	if err != nil {
		t.Fatal(err)
	}
	b := make([]byte, chunkSize)
	for i := 0; i < int(chunkNum); i++ {
		_, err = myFile.Seek(int64(i*chunkSize), 0)
		if chunkSize > fileinfo.Size()-int64(i*chunkSize) {
			b = make([]byte, fileinfo.Size()-int64(i*chunkSize))
		}
		myFile.Read(b)                                                                             //把上传的文件放入 b 中，再把这个b保存
		f, err := os.OpenFile("./"+strconv.Itoa(i)+".chunk", os.O_CREATE|os.O_WRONLY, os.ModePerm) //新建chunk文件
		defer f.Close()
		if err != nil {
			t.Fatal(err)
		}
		f.Write(b) //数据流写入chunk
	}
}

// 分片文件合并
func TestMergeChunkFile(t *testing.T) {
	myFile, err := os.OpenFile("test2.jpg", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm) //新文件
	defer myFile.Close()
	if err != nil {
		t.Fatal(err)
	}

	//分片个数
	fileinfo, err := os.Stat("../img/test.jpg") //源文件
	if err != nil {
		t.Fatal(err)
	}
	chunkNum := math.Ceil(float64(fileinfo.Size()) / float64(chunkSize))

	//读取文件
	for i := 0; i < int(chunkNum); i++ {
		f, err := os.OpenFile("./"+strconv.Itoa(i)+".chunk", os.O_RDONLY, 0666)
		defer f.Close()
		if err != nil {
			t.Fatal(err)
		}
		b, err := io.ReadAll(f)
		if err != nil {
			t.Fatal(err)
		}
		myFile.Write(b)
	}
}

// 文件一致性校验
func TestCheck(t *testing.T) {
	// 原文件
	file1, err := os.OpenFile("test.flac", os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	b1, err := io.ReadAll(file1)
	if err != nil {
		t.Fatal(err)
	}

	// 新文件
	file2, err := os.OpenFile("test2.flac", os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	b2, err := io.ReadAll(file2)
	if err != nil {
		t.Fatal(err)
	}

	s1 := fmt.Sprintf("%x", md5.Sum(b1))
	s2 := fmt.Sprintf("%x", md5.Sum(b2))
	fmt.Println(s1)
	fmt.Println(s2)
	if s1 != s2 {
		t.Fatal("文件不一致")
	}
}
