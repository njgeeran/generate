package service

import (
	"fmt"
	"github.com/njgeeran/core/utils"
	"lowcode/generate/model"
	"os"
	"strings"
	"text/template"
)

func generate_routers(msg chan string,root_path string,p *model.Project) error {
	router_path := root_path+"/router"
	router_file_path := router_path+"/router.go"

	if err := generate_router_path(msg,router_path);err != nil {
		return err
	}

	if err := generate_router_file(msg,router_file_path,p);err != nil {
		return err
	}

	if err := generate_api_handers(msg,root_path,p);err != nil {
		return err
	}

	return nil
}

func generate_router_path(msg chan string,router_path string) error {
	SendGenerateMsg(msg,"生成路由目录:"+router_path)
	if err := utils.MkdirIfNotExist(router_path);err != nil{
		SendGenerateErrMsg(msg,"路由目录创建失败",err)
		return err
	}
	return nil
}

var router_temp = `package router

func InitRouter(router *gin.RouterGroup)  {
{{.apis}}
}
`
func generate_router_file(msg chan string,router_file_path string,p *model.Project) error {
	SendGenerateMsg(msg,"开始生成路由:"+router_file_path)
	apis := []string{}
	for _,t := range p.Routers {
		method := ""
		switch t.Method {
		case model.ApiMethod_GET:
			method = "GET"
		case model.ApiMethod_POST:
			method = "POST"
		case model.ApiMethod_PUT:
			method = "PUT"
		}
		api := fmt.Sprintf("\trouter.%s(\"%s\",api.%s)",method,t.Path,t.Name)
		apis = append(apis, api)
	}
	f,err := os.Create(router_file_path)
	if err != nil {
		SendGenerateErrMsg(msg,"创建路由文件失败",err)
		return err
	}
	tmpl := template.Must(template.New("router").Parse(router_temp))
	tmpl.Execute(f, map[string]string{
		"apis":strings.Join(apis,"\n"),
	})
	f.Close()
	return nil
}