package service

import (
	"fmt"
	"github.com/njgeeran/core/utils"
	"lowcode/generate/model"
	"os"
	"strings"
	"text/template"
)

func generate_api_handers(msg chan string,root_path string,p *model.Project) error {
	SendGenerateMsg(msg,"生成api hander")
	api_path := root_path+"/api"
	api_hander_path := api_path+"/hander"

	if err := generate_api_path(msg,api_path);err != nil {
		return err
	}
	if err := generate_api_hander_path(msg,api_hander_path);err != nil {
		return err
	}

	if err := generate_apis(msg,api_path,api_hander_path,p);err != nil {
		return err
	}

	return nil
}

func generate_api_path(msg chan string,api_path string) error {
	SendGenerateMsg(msg,"创建api目录:"+api_path)
	if err := utils.MkdirIfNotExist(api_path);err != nil{
		SendGenerateErrMsg(msg,"api目录创建失败",err)
		return err
	}
	return nil
}
func generate_api_hander_path(msg chan string,api_hander_path string) error {
	SendGenerateMsg(msg,"创建api hander目录:"+api_hander_path)
	if err := utils.MkdirIfNotExist(api_hander_path);err != nil{
		SendGenerateErrMsg(msg,"api hander目录创建失败",err)
		return err
	}
	return nil
}
func generate_apis(msg chan string,api_path,api_hander_path string,p *model.Project) error {
	for _,t := range p.Routers {
		if err := generate_api_file(msg,api_path,p,&t);err != nil {
			return err
		}
		if err := generate_api_hander_file(msg,api_hander_path,p,&t);err != nil {
			return err
		}
	}
	return nil
}

var api_temp = `package api

func {{.name}}(c *gin.Context)  {
{{.url_para}}
{{.body_para}}

	h := hander.New{{.name}}Hander(c)
	{{.func_resp_str}} := h.{{.name}}({{.hander_para}})
	if err != nil {
		httpx.FailWithMessage(c,err.Error())
		return
	}
	httpx.OkWithData(c,{{.resp_val}})
}
`
func generate_api_file(msg chan string,api_path string,p *model.Project,r *model.Router) error {
	//生成api文件
	name := r.Name
	url_para := get_api_file_url_para(r)
	body_para := get_api_file_body_para(r)
	func_resp_str := utils.IF(r.Hander.ReturnVal == nil, "err", "data,err").(string)
	resp_val := utils.IF(r.Hander.ReturnVal == nil, "nil", "data").(string)
	hander_para := get_api_file_hander_para(r)

	f, err := os.Create(api_path + "/" + utils.CamelCaseToUnderline(name) + ".go")
	if err != nil {
		return err
	}
	tmpl := template.Must(template.New("api").Parse(api_temp))
	tmpl.Execute(f, map[string]string{
		"name":          name,
		"url_para":      url_para,
		"body_para":     body_para,
		"hander_para":   hander_para,
		"func_resp_str": func_resp_str,
		"resp_val":      resp_val,
	})
	f.Close()
	return nil
}
func get_api_file_url_para(r *model.Router) string {
	if len(r.Hander.UrlParaModel) <= 0 {
		return ""
	}
	url_para_str := []string{}
	for _,t := range r.Hander.UrlParaModel {
		str := ""
		switch t.FieldType {
		case "int":
			str = fmt.Sprintf("\t%s,_ := strconv.Atoi(c.DefaultQuery(\"%s\",\"%s\"))",t.FieldName,t.FieldName,t.DefaultValue)
		case "string":
			str = fmt.Sprintf("\t%s := c.DefaultQuery(\"%s\",\"%s\")",t.FieldName,t.FieldName,t.DefaultValue)

		}
		if len(t.Validators) > 0 {
			str += "\n\t"+t.GenerateValidators()
		}
		url_para_str = append(url_para_str, str)
	}
	return strings.Join(url_para_str,"\n")
}
func get_api_file_body_para(r *model.Router) string {
	if r.Hander.BodyParaModel == nil {
		return ""
	}
	body_para_str := ""
	body_para_str = fmt.Sprintf("\treqStruct := &model.%s{}\n",r.Hander.BodyParaModel.JoinModel.Name)
	switch r.Hander.BodyParaModel.ContentType {
	case model.BodyParaContentType_JSON:
		body_para_str += "\tc.ShouldBindJSON(reqStruct)\n"
	case model.BodyParaContentType_FORMDATA:
		body_para_str += "\tc.ShouldBind(reqStruct)\n"
	}
	//validators
	body_para_str += r.Hander.BodyParaModel.Validators.GenerateGo()
	return body_para_str
}
func get_api_file_hander_para(r *model.Router) string {
	hander_para := []string{}
	for _,t := range r.Hander.UrlParaModel {
		hander_para = append(hander_para,t.FieldName)
	}
	if r.Hander.BodyParaModel != nil {
		hander_para = append(hander_para, "reqStruct")
	}
	return strings.Join(hander_para,",")
}


var api_hander_temp = `package hander

type {{.name}}Hander struct {
	log 	*log.Loger
	c    	*gin.Context
}

func New{{.name}}Hander(c *gin.Context) {{.name}}Hander {
	return {{.name}}Hander{
		c:c,
		log:log.GetLoger(),
	}
}
func (h *{{.name}}Hander) {{.name}}({{.hand_para_str}}) ({{.para_resp_val}}) {
{{.hander_code}}
}

`
func generate_api_hander_file(msg chan string,api_hander_path string,p *model.Project,r *model.Router) error {
	name := r.Name
	hand_para_str := get_api_hander_file_hand_para_str(r)
	para_resp_val := get_api_hander_file_para_resp_val(r)

	f,err := os.Create(api_hander_path+"/"+utils.CamelCaseToUnderline(name)+"_hander.go")
	if err != nil {
		return err
	}
	tmpl := template.Must(template.New("hander").Parse(api_hander_temp))
	tmpl.Execute(f, map[string]string{
		"name":name,
		"hand_para_str":hand_para_str,
		"hander_code":r.Hander.HanderCode,
		"para_resp_val":para_resp_val,
	})
	f.Close()
	return nil
}
func get_api_hander_file_hand_para_str(r *model.Router) string {
	hander_para := []string{}
	for _,t := range r.Hander.UrlParaModel {
		hander_para = append(hander_para,t.FieldName+" "+t.FieldType)
	}
	if r.Hander.BodyParaModel != nil {
		hander_para = append(hander_para, fmt.Sprintf("reqStruct *model.%s",r.Hander.BodyParaModel.JoinModel.Name))
	}
	return strings.Join(hander_para,",")
}
func get_api_hander_file_para_resp_val(r *model.Router) string {
	if r.Hander.ReturnVal == nil {
		return "err error"
	}
	return fmt.Sprintf("data *%s,err error",r.Hander.ReturnVal.FieldType)
}