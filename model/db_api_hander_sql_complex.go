package model

import "github.com/njgeeran/core/utils"

type ApiHanderSqlComplex struct {
	utils.BaseModel
	UrlPara				string			`json:"url_para"`	//url参数
	BodyPara			string			`json:"body_para"`	//model id|content_type
	ReturnVal			string			`json:"return_val"`	//名称(model id)|类型|
}