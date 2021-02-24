package service

import (
	"fmt"
	"github.com/njgeeran/generate/model"
)

func generate_gomod(msg chan string,go_path,root_path string,p *model.Project) error {
	if err := generate_init_gomod(msg,go_path,root_path,p);err != nil {
		return err
	}
	if err := generate_init_get_module(msg,go_path,root_path,p);err != nil {
		return err
	}
	if err := generate_install_module(msg,go_path,root_path,p);err != nil {
		return err
	}
	if err := generate_gomod_tidy(msg,go_path,root_path,p);err != nil {
		return err
	}
	return nil
}
func generate_init_gomod(msg chan string,go_path,root_path string,p *model.Project) error {
	SendGenerateMsg(msg,"生成gomod:"+root_path)
	result,err := ExecCommandGo(root_path,"",go_path,"go","mod","init",p.Name)
	SendGenerateMsg(msg,"生成gomod输出："+result)
	if err != nil {
		SendGenerateErrMsg(msg,"gomod生成失败",err)
		return err
	}
	return nil
}
func generate_init_get_module(msg chan string,go_path,root_path string,p *model.Project) error {
	for _,t := range p.Modules {
		fmt.Println(p.Modules)
		m_name := t.Path+"@"+t.Version
		SendGenerateMsg(msg,fmt.Sprintf("获取依赖模块[%s]",m_name))
		result,err := ExecCommandGo(root_path,"",go_path,"go","get","-v",m_name)
		SendGenerateMsg(msg,fmt.Sprintf("获取依赖模块[%s]输出："+result,m_name))
		if err != nil {
			SendGenerateErrMsg(msg,fmt.Sprintf("获取依赖模块[%s]失败",m_name),err)
			return err
		}
	}

	return nil
}
func generate_install_module(msg chan string,go_path,root_path string,p *model.Project) error {
	SendGenerateMsg(msg,"插入依赖")
	result,err := ExecCommandGo(root_path,"",go_path,"goimports","-v","-w",root_path)
	SendGenerateMsg(msg,"插入依赖输出："+result)
	if err != nil {
		SendGenerateErrMsg(msg,"插入依赖失败",err)
		return err
	}
	return nil
}
func generate_gomod_tidy(msg chan string,go_path,root_path string,p *model.Project) error {
	SendGenerateMsg(msg,"整理gomode")
	result,err := ExecCommandGo(root_path,"",go_path,"go","mod","tidy")
	SendGenerateMsg(msg,"整理gomode输出："+result)
	if err != nil {
		SendGenerateErrMsg(msg,"整理gomode失败",err)
		return err
	}
	return nil
}