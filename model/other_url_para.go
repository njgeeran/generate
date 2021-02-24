package model

import (
	"fmt"
	"strings"
)

type UrlPara struct {
	FieldName 		string		`json:"field_name"`			//字段名称
	FieldType 		string		`json:"field_type"`			//字段类型
	DefaultValue 	string		`json:"default_value"`		//默认值
	Validators		[]string	`json:"validators"`			//验证规则
}
func (u UrlPara)GenerateValidators() string {
	roles := []string{}
	t := u.Validators
	for i := 0; i < len(t); i++ {
		t[i] = fmt.Sprintf("\"%s\"", t[i])
	}
	role := fmt.Sprintf("\t\t\"%s\":{%s},", u.FieldName, strings.Join(t, ","))

	roles = append(roles, role)
	if len(roles) <= 0 {
		return ""
	}
	str := fmt.Sprintf("\tvar "+u.FieldName+"Validators = map[string][]string{\n%s\n\t}\n", strings.Join(roles, "\n"))
	str += "\tif err := utils.VerifyField(\""+u.FieldName+"\","+u.FieldName+","+u.FieldName+"Validators);err != nil{\n\t\thttpx.FailWithMessage(c,err.Error())\n\t\treturn\n\t}\n"
	return str
}