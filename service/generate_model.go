package service

import (
	"fmt"
	"github.com/njgeeran/core/utils"
	"io"
	"lowcode/generate/model"
	"os"
	"strings"
	"text/template"
)

func generate_models(msg chan string,root_path string,p *model.Project) error {
	SendGenerateMsg(msg,"开始生成模型")
	m_path := root_path+"/model"

	//生成模型路径
	if err := generate_models_path(msg,m_path);err != nil {
		return err
	}

	//生成数据源操作(db_name)
	if err := generate_models_db_name(msg,m_path,p);err != nil {
		return err
	}

	//生成模型
	for _,m := range p.Models {
		if err := generate_model_file(msg,&m,m_path);err != nil{
			return err
		}
	}
	return nil
}
func generate_models_path(msg chan string,m_path string) error {
	//生成model目录
	SendGenerateMsg(msg,"开始生成模型目录:"+m_path)
	if err := utils.MkdirIfNotExist(m_path);err != nil{
		SendGenerateErrMsg(msg,"生成模型路径失败",err)
		return err
	}
	return nil
}
func generate_models_db_name(msg chan string,m_path string,p *model.Project) error {
	db_name_file_path := m_path + "/db_name.go"
	SendGenerateMsg(msg, "开始生成数据库操作模型:"+db_name_file_path)
	db_name_strs := []string{}
	for _, t := range p.DataSources {
		name := utils.UnderlinToCamelCase(t.Name)
		db_name_str := fmt.Sprintf("const Db%s = \"%s\"\n", name, t.Name)
		db_name_str += fmt.Sprintf("func GetDb%s() *gorm.DB {\n\treturn orm.GetOrm().Get(Db%s)\n}", name, name)
		db_name_strs = append(db_name_strs, db_name_str)
	}
	f, err := os.Create(db_name_file_path)
	if err != nil {
		SendGenerateErrMsg(msg,"数据库操作模型文件创建失败",err)
		return err
	}
	f.Write([]byte("package model\n\n"))
	f.Write([]byte(strings.Join(db_name_strs, "\n\n")))
	f.Close()
	return nil
}

func generate_model_file(msg chan string,m *model.Model,m_path string) error {
	file_prefix := ""
	switch m.Type {
	case model.ModelType_Db:
		file_prefix = "db_"
	case model.ModelType_ReqStruct:
		file_prefix = "req_"
	case model.ModelType_RespStruct:
		file_prefix = "resp_"
	case model.ModelType_Other:
		file_prefix = "other_"
	}
	m_go_path := m_path+"/"+file_prefix+utils.CamelCaseToUnderline(m.Name)+".go"
	SendGenerateMsg(msg, "开始生成模型:"+m_go_path)

	f,err := os.Create(m_go_path)
	if err != nil {
		SendGenerateErrMsg(msg,"模型文件创建失败",err)
		return err
	}
	GenerateModel(m,f)
	f.Close()
	return nil
}
var model_temp = `package model

{{.desc}}
type {{.ModelName}} struct {
{{.fields}}
}

{{.enums}}
`
func GenerateModel(m *model.Model,wr io.Writer) {
	fields := []string{}
	enums := ""

	if m.IsUseBaseModel {
		fields = append(fields,"\tutils.BaseModel")
	}
	for _,t := range m.Fields {
		fields = append(fields,fmt.Sprintf("\t%s",GenerateField(&t)))
		if t.Type == model.FieldType_ENUM {
			enums += fmt.Sprintf("\n%s\n",GenerateEnum(t.Other))
		}
	}

	tmpl := template.Must(template.New("model").Parse(model_temp))
	tmpl.Execute(wr, map[string]string{
		"fields":strings.Join(fields,"\n"),
		"desc":utils.IF(m.Desc == "",m.Desc,"//"+m.Desc).(string),
		"ModelName":m.Name,
		"enums":enums,
	})

	return
}