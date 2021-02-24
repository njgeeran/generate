package model

type ApiHander struct {
	UrlParaModel 		[]UrlPara
	BodyParaModel 		*BodyPara
	ReturnVal			*ReturnVal
	HanderCode			string			`json:"hander_code"`
}

