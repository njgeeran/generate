package service

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func ExecCommandGo(dir,path,go_path string,command string,arg ...string) (string,error) {
	goroot := os.Getenv("GOROOT")
	env := []string{
		"GO111MODULE=on",
		"GOPROXY=https://mirrors.aliyun.com/goproxy/",
		"GOPATH="+go_path,
		"GOROOT="+goroot,
		"GOBIN="+goroot+"/bin",
		"path="+goroot+"/bin",
		fmt.Sprintf("GOCACHE=%s/go-build",os.TempDir()),
		"GOTMPDIR="+os.TempDir(),
	}
	return ExecCommand(dir,path,env,command,arg...)
}
func ExecCommand(dir string,path string,env []string,command string,arg ...string) (string,error) {
	cmd := exec.Command(command, arg...)
	if dir != "" {
		cmd.Dir = dir
	}
	if path != "" {
		cmd.Path = path
	}
	cmd.Env = env

	outinfo := &bytes.Buffer{}
	cmd.Stdout = outinfo
	cmd.Stderr = outinfo

	if err := cmd.Run(); err != nil {   // 运行命令
		return "",errors.New(err.Error()+":"+outinfo.String())
	}
	return outinfo.String(),nil
}
