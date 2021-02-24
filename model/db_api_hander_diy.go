package model

import (
	"encoding/json"
	"github.com/njgeeran/core/utils"
)

type ApiHanderDiy struct {
	utils.BaseModel
	UrlPara				string			`json:"url_para"`	//url参数
	BodyPara			string			`json:"body_para"`	//model id|content_type
	ReturnVal			string			`json:"return_val"`	//名称(model id)|类型|
	HanderCode			string			`json:"hander_code"`
}

func (h *ApiHanderDiy)ToApiHander(p *Project) (*ApiHander,error) {
	ah := &ApiHander{}
	if h.UrlPara != "" {
		if err := json.Unmarshal([]byte(h.UrlPara),&ah.UrlParaModel);err != nil{
			return nil,err
		}
	}

	if h.BodyPara != "" {
		if err := json.Unmarshal([]byte(h.BodyPara),&ah.BodyParaModel);err != nil{
			return nil,err
		}
		for _,t := range p.Models {
			if t.ID == ah.BodyParaModel.ModelId {
				ah.BodyParaModel.JoinModel = t
				break
			}
		}
	}

	if h.ReturnVal != "" {
		if err := json.Unmarshal([]byte(h.ReturnVal),&ah.ReturnVal);err != nil{
			return nil,err
		}
	}
	ah.HanderCode = h.HanderCode
	return ah,nil
}