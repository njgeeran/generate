package service

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"github.com/njgeeran/generate/model"
	"testing"
)

var m1 = &model.Model{
	Name:"User",
	Desc:"用户表",
	Fields:[]model.Field{
		model.Field{
			Name:"Id",
			Type:model.FieldType_INT,
		},
		model.Field{
			Name:"Name",
			Type:model.FieldType_STRING,
			Desc:"用户名称",
		},
		model.Field{
			Name:"Type",
			Type:model.FieldType_ENUM,
			Other:"UserType|int|Admin,Teacher,Student",
		},
	},
	IsUseBaseModel:true,
}

func TestGenerateModel(t *testing.T) {
	wr := bytes.NewBuffer([]byte{})
	GenerateModel(m1,wr)
	result,_ := ioutil.ReadAll(wr)
	fmt.Println(result)
}
