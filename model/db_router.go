package model

import "github.com/njgeeran/core/utils"

//todo 中间件
//todo 层级结构
type Router struct {
	utils.BaseModel
	FromProject 	int 		`json:"from_project"` //属于项目
	Name		string	`json:"name"`		//名称
	Path		string 	`json:"path"`		//地址
	Method		ApiMethod	`json:"method"`	//方法
	HanderId	int		`json:"hander_id"`	//处理器id
	HanderType	ApiHanderType		`json:"hander_type"`	//处理器类型

	Hander 		*ApiHander
}

type ApiMethod int
const (
	_ApiMethod = iota
	ApiMethod_GET
	ApiMethod_POST
	ApiMethod_PUT
)

type ApiHanderType int
const (
	_ApiHanderType = iota
	ApiHanderType_EasySql
	ApiHanderType_ComplexSql
	ApiHanderType_DIY
)