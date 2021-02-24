package model

import "github.com/njgeeran/core/utils"

type Module struct {
	utils.BaseModel
	FromProject int			`json:"from_project"`
	Path 		string		`json:"path"`		//路径
	Version 	string		`json:"version"`	//版本
}
