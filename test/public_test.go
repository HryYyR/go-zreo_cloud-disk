package test

import (
	"cloud_disk/go-zreo_cloud-disk/core/define"
	"cloud_disk/go-zreo_cloud-disk/core/helper"
	"crypto/tls"
	"fmt"
	"net/smtp"
	"testing"

	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
)

func TestSendEmail(t *testing.T) {
	e := email.NewEmail()
	e.From = "zero-cloud-disk <2452719312@qq.com>"
	e.To = []string{"2269967055@qq.com"}
	e.Subject = "验证码发送测试" //标题
	// e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("<h1>验证码: 123456</h1>") //内容
	// e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "test@gmail.com", "password123", "smtp.gmail.com"))
	err := e.SendWithTLS("smtp.qq.com:465", smtp.PlainAuth("", "2452719312@qq.com", define.EmailPassword, "smtp.qq.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRandCode(t *testing.T) {
	code := helper.RandCode()
	fmt.Println(code)
}

// token加密
func TestGenerateUUID(t *testing.T) {
	u1 := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", u1)
}

// token解密
func TestDecryptToekn(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6NiwiSWRlbnRpdHkiOiJjMmY0YjYyNi04NTJjLTRiMzYtOTZhMy1lYzA1OTlkZWU5NGEiLCJOYW1lIjoiSHl5eWgiLCJleHAiOjE2OTA3NTk5NDN9.1MeAJpy6xYi25T_u90_c3OMUTSAGjiX2pdkCLWzjeww"
	detoken, err := helper.DecryptToekn(token)
	fmt.Printf("%v\n", detoken)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", detoken)
}
