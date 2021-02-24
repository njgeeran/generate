package service

import (
	"fmt"
	"testing"
)

func TestGenerateMaingo(t *testing.T) {
	InitTest()
	p,_ := GetProjectData(GetOrm(),1)
	msg := make(chan string,200)
	result,err := GenerateMaingo(msg,p)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)
}
