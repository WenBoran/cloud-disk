package test

import (
	"bytes"
	"clou-disk/models"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"xorm.io/xorm"
)

func TestXorm(t *testing.T) {

	var engine *xorm.Engine
	var err error
	engine, err = xorm.NewEngine("mysql", "root:123@/cloud-disk?charset=utf8")
	if err != nil {
		t.Fatal(err)
	}
	data := make([]*models.UserBasic, 0)
	err = engine.Find(&data)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(data)
	b, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	dst := new(bytes.Buffer)
	json.Indent(dst, b, "", "  ")

	fmt.Println(dst.String())
}
