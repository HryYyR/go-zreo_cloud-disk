package test

import (
	"bytes"
	"cloud_disk/go-zreo_cloud-disk/core/models"
	"encoding/json"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func TestXormTest(t *testing.T) {
	engine, err := xorm.NewEngine("mysql", "root:root@/cloud-disk?charset=utf8mb4")
	if err != nil {
		t.Fatal(err)

		fmt.Println("connect mysql failed", err.Error())
	}
	data := make([]*models.UserBasic, 0)
	err = engine.Find(&data)
	if err != nil {
		t.Fatal(err)
	}

	b, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	dst := new(bytes.Buffer)
	err = json.Indent(dst, b, "", " ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dst.String())
}
