package service

import (
	"fmt"
	"lowcode/generate/model"
	"testing"
)

func TestGenerateEnum(t *testing.T) {
	str := "UserType|string|Admin,Teacher,Student"
	result := GenerateEnum(str)
	fmt.Println(result)
}

func TestGenerateField(t *testing.T) {
	f := &model.Field{
		Name:"UserType",
		Type:model.FieldType_ENUM,
		Other:"UserType|string|Admin,Teacher,Student",
		Desc:"用户类型",
	}
	result := GenerateField(f)
	fmt.Println(result)
	//---------------
	f.Name = "UserName"
	f.Type = model.FieldType_STRING
	f.Other = ""
	f.Desc = "用户名称"
	result = GenerateField(f)
	fmt.Println(result)
	//---------------
	f.Name = "TestModel"
	f.Type = model.FieldType_Model
	f.Other = "[]Model"
	f.Desc = "测试Model"
	f.GormSetting = "size:255;AUTO_INCREMENT"
	result = GenerateField(f)
	fmt.Println(result)
}
