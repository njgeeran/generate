package service

import (
	"fmt"
	"github.com/njgeeran/core/utils"
	"lowcode/generate/model"

	"strings"
)

func GenerateField(f *model.Field) string {
	types := ""
	switch f.Type {
	case model.FieldType_INT:
		types = "int"
	case model.FieldType_STRING:
		types = "string"
	case model.FieldType_TIME:
		types = "time.Time"
	case model.FieldType_Bool:
		types = "bool"
	case model.FieldType_ENUM:
		types = strings.Split(f.Other,"|")[0]
	case model.FieldType_Model:
		types = f.Other
	}

	settings := []string{}
	json_name := f.JsonName
	if json_name == "" {
		json_name = utils.CamelCaseToUnderline(f.Name)
	}
	settings = append(settings, fmt.Sprintf("json:\"%s\"",json_name))
	if f.GormSetting != "" {
		settings = append(settings,fmt.Sprintf("gorm:\"%s\"",f.GormSetting))
	}
	setting_str := fmt.Sprintf("`%s`",strings.Join(settings," "))


	return fmt.Sprintf("%s\t%s\t%s\t%s",
		f.Name,types,setting_str,
		utils.IF(f.Desc == "",f.Desc,"//"+f.Desc),
	)
}

//Name(英文，首字母大写)|类型(int,string)|value(,分割)
func GenerateEnum(str string) string {
	strs := strings.Split(str,"|")
	if len(strs) < 3 {
		return ""
	}
	var (
		name = strs[0]
		types = strs[1]
		value = strs[2]
	)
	types = strings.ToLower(types)
	if types != "int" && types != "string" {
		return ""
	}

	result := ""
	result += fmt.Sprintf("type %s %s\n",name,types)

	val := ""
	if types == "int" {
		val += fmt.Sprintf("\t_%s 		=	iota\n",name)
	}
	vals := strings.Split(value,",")
	for _,t := range vals {
		switch types {
		case "int":
			val += fmt.Sprintf("\t%s_%s \n",name,t)
			break
		case "string":
			val += fmt.Sprintf("\t%s_%s = \"%s\" \n",name,t,t)
			break
		}

	}

	result += fmt.Sprintf("const ( \n %s )",val)
	return result
}
