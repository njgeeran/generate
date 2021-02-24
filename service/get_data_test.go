package service

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetProjectData(t *testing.T) {
	InitTest()
	p,err := GetProjectData(GetOrm(),1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	p_json,_ := json.Marshal(p)
	fmt.Println(string(p_json))
}
