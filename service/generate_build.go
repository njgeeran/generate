package service

import (
	"github.com/njgeeran/generate/model"
	"runtime"
)

func generate_build(msg chan string,go_path,root_path string,p *model.Project) error {
	SendGenerateMsg(msg,"编译源码，生成可执行文件")
	out := p.Name
	if runtime.GOOS == "windows" {
		out += "-win.exe"
	}
	result,err := ExecCommandGo(root_path,"",go_path,"go","build","-x","-o",out,"main.go")
	SendGenerateMsg(msg,"编译源码输出："+result)
	if err != nil {
		SendGenerateErrMsg(msg,"编译源码失败",err)
		return err
	}
	return nil
}
