package model

import "github.com/njgeeran/core/utils"

type Project struct {
	utils.BaseModel
	Name string `json:"name"` //项目名称 【英文,首字母大写】
	Desc string `json:"desc"` //项目说明

	//---------------------------------------------
	Models      []Model
	DataSources []DataSource
	Modules 	[]Module
	Routers 	[]Router
}