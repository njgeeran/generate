package service

import (
	"fmt"
	"testing"
)

func TestExecCommandGo(t *testing.T) {
	result,err := ExecCommandGo("C:/Users/ADMINI~1/AppData/Local/Temp/low_code_generate/src/1","","C:/Users/ADMINI~1/AppData/Local/Temp/low_code_generate","go","env")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)
}
