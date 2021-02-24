package service

import (
	"fmt"
	"testing"
)

func TestGenerateProject(t *testing.T) {
	InitTest()
	p,err := GetProjectData(GetOrm(),1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	msg := make(chan string,200)
	over := make(chan bool,1)
	GenerateProject(msg,over,p)
	for {
		select {
		case m := <-msg:
				fmt.Println(m)
		case isok := <-over:
			if isok {
				fmt.Println("生成成功")
			}else {
				fmt.Println("生成失败")
			}
			return
		}
	}
}
