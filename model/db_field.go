package model

import "github.com/njgeeran/core/utils"

type Field struct {
	utils.BaseModel
	FromModel 	int 	`json:"from_model"` //属于模型

	Name        string    `json:"name"`         //名称【英文,首字母大写】
	Desc        string    `json:"desc"`         //注释
	Type        FieldType `json:"type"`         //类型
	Other       string    `json:"other"`        //枚举设置Enum[Name(英文，首字母大写)|类型(int,string)|value(,分割)]、Model[model名称]
	JsonName    string    `json:"json_name"`    //Json名称
	GormSetting string    `json:"gorm_setting"` //gorm设置
}
type FieldType int
const (
	_FieldType = iota
	FieldType_INT
	FieldType_STRING
	FieldType_TIME
	FieldType_ENUM
	FieldType_Model
	FieldType_Bool
)
