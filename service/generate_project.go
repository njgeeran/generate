package service

import (
	"github.com/google/uuid"
	"github.com/njgeeran/core/utils"
	"lowcode/generate/model"
	"os"
	"path/filepath"
	"strconv"
)

func GenerateProject(msg chan string,over chan bool,p *model.Project) string {
	output := "/output/"+uuid.New().String()+".tar.gz"
	go generate_project(msg,over,p,output)
	return output
}

func generate_project(msg chan string,over chan bool,p *model.Project,output string)  {
	var (
		path = os.TempDir()+"/low_code_generate"
		root_path = path+"/src/"+strconv.Itoa(p.ID)
		isok = true
	)

	defer func() {
		os.RemoveAll(root_path)
		over <- isok
	}()

	SendGenerateMsg(msg,"开始生成项目")

	//生成主目录并删除主目录下所有文件
	if err := generate_root_path(msg,root_path);err != nil {
		isok = false
		return
	}

	//生成main入口文件
	if err := generate_maingo(msg,root_path,p);err != nil {
		isok = false
		return
	}

	//生成配置文件
	if err := generate_config(msg,root_path,p);err != nil {
		isok = false
		return
	}

	//生成模型
	if err := generate_models(msg,root_path,p);err != nil{
		isok = false
		return
	}

	//生成路由和接口
	if err := generate_routers(msg,root_path,p);err != nil {
		isok = false
		return
	}

	//生成gomod和插入依赖
	if err := generate_gomod(msg,path,root_path,p);err != nil {
		isok = false
		return
	}

	//生成可执行文件
	if err := generate_build(msg,path,root_path,p);err != nil {
		isok = false
		return
	}

	//处理输出文件
	if err := compress_output(msg,root_path,output);err != nil {
		isok = false
		return
	}
}
//生成主目录
func generate_root_path(msg chan string,root_path string) error {
	SendGenerateMsg(msg,"开始生成主目录:"+root_path)
	os.RemoveAll(root_path)
	//生成主目录
	if err := utils.MkdirIfNotExist(root_path);err != nil{
		SendGenerateErrMsg(msg,"主目录生成失败:",err)
		return err
	}
	return nil
}
func compress_output(msg chan string,source_path,output_path string) error {
	output_path,err := filepath.Abs(output_path)
	if err != nil {
		SendGenerateErrMsg(msg,"处理输出文件失败",err)
		return err
	}
	if err := utils.CompressDirNotIn(source_path,output_path);err != nil {
		SendGenerateErrMsg(msg,"处理输出文件失败",err)
		return err
	}
	return nil
}