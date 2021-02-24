package model

import "github.com/njgeeran/core/utils"

type Model struct {
	utils.BaseModel
	FromProject 	int 		`json:"from_project"` //属于项目
	FromDS      	int 		`json:"from_ds"`      //属于数据源

	Name           	string    	`json:"name"`              //名称【英文,首字母大写】
	Type           	ModelType 	`json:"type"`              //类型
	Desc           	string    	`json:"desc"`              //注释
	IsUseBaseModel 	bool      	`json:"is_use_base_model"` //是否使用BaseModel

	IsAutoMigrate 	bool `json:"is_auto_migrate"` //是否自动迁移

	Fields []Field
}
type ModelType int
const (
	_ModelType = iota
	ModelType_Db
	ModelType_ReqStruct
	ModelType_RespStruct
	ModelType_Other
)