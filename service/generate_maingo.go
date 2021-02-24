package service

import (
	"bytes"
	"fmt"
	"github.com/njgeeran/core/utils"
	"github.com/njgeeran/generate/model"
	"os"
	"strings"
	"text/template"
)

//生成main.go(程序入口)文件
var maingo_temp = `package main

func init()  {
	{{.init}}
}

{{.init_func}}

func main()  {
	init_server()
}

//加载http监听
func init_server() {
	//开启gin
	r := core.InitGin()
	//挂载路由
	//api
	ApiGroup := r.Group("")
	router.InitRouter(ApiGroup)

	log.GetLoger().Error(r.Run(":"+conf.GetConf().SysSetting.HttpAddr))
}
`

func GenerateMaingo(msg chan string,p *model.Project) (string,error) {
	//init func
	inits := []string{}
	init_funcs := []string{}
	//init orm
	init_orm_funcs := ""
	for i,t := range p.DataSources {
		init_orm_func := ""
		is_need := false
		if t.SingularTable {
			is_need = true

			init_orm_func += fmt.Sprintf("\tdb%d.SingularTable(true)\n",i)
		}
		//自动迁移
		auto_migrates := []string{}
		for _,tt := range t.Models {
			if tt.IsAutoMigrate {
				auto_migrates = append(auto_migrates, "model."+tt.Name+"{}")
			}
		}
		if len(auto_migrates) > 0 {
			is_need = true
			init_orm_func += fmt.Sprintf("\tdb%d.AutoMigrate(%s)\n",i,strings.Join(auto_migrates,","))
		}
		//结束
		if is_need {
			init_orm_func = fmt.Sprintf("\tdb%d := model.GetDb%s()\n",i,utils.UnderlinToCamelCase(t.Name))+init_orm_func
			init_orm_funcs += init_orm_func
		}
	}
	if init_orm_funcs != "" {
		inits = append(inits, "init_db_tables()")
		init_funcs = append(init_funcs, fmt.Sprintf("// 注册数据库表\nfunc init_db_tables() {\n%s}",init_orm_funcs))
	}

	w := bytes.NewBufferString("")
	tmpl := template.Must(template.New("main").Parse(maingo_temp))
	tmpl.Execute(w, map[string]string{
		"init":strings.Join(inits,"\n"),
		"init_func":strings.Join(init_funcs,"\n"),
	})
	return w.String(),nil
}
func generate_maingo(msg chan string,root_path string,p *model.Project) error {
	main_file_path := root_path+"/main.go"
	SendGenerateMsg(msg,"开始生成main入口文件:"+main_file_path)

	f_str,err := GenerateMaingo(msg,p)
	if err != nil {
		SendGenerateErrMsg(msg,"main入口文件内容生成失败:",err)
		return err
	}

	f,err := os.Create(main_file_path)
	if err != nil {
		SendGenerateErrMsg(msg,"main入口文件创建失败:",err)
		return err
	}
	if _,err := f.WriteString(f_str);err != nil{
		SendGenerateErrMsg(msg,"main入口文件写入失败:",err)
	}
	f.Close()
	return nil
}