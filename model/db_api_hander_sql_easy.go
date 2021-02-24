package model

import (
	"encoding/json"
	"errors"
	"github.com/njgeeran/core/utils"
)

type ApiHanderSqlEasy struct {
	utils.BaseModel
	ModelId    	int                 	`json:"model_id"`
	Type    	ApiHanderSqlEasyType 	`json:"type"`                     //简易sql类型
	Setting 	string               	`json:"setting" gorm:"type:text"` //配置

	JoinModel	Model
}
type ApiHanderSqlEasyType int
const (
	_ApiHanderSqlEasyType = iota
	ApiHanderSqlEasyType_ADD
	ApiHanderSqlEasyType_DELETE
	ApiHanderSqlEasyType_UPDATE
	ApiHanderSqlEasyType_SEARCHFIRST
	ApiHanderSqlEasyType_SEARCHLIST
)

func (h *ApiHanderSqlEasy)ToApiHander(p *Project) (*ApiHander,error) {
	for _,t := range p.Models {
		if t.ID == h.ModelId {
			h.JoinModel = t
		}
	}
	switch h.Type {
	case ApiHanderSqlEasyType_ADD:
		{
			hander := &ApiHanderSqlEasySettingAdd{}
			if h.Setting != "" {
				if err := json.Unmarshal([]byte(h.Setting), hander); err != nil {
					return nil, err
				}
			}
			return hander.ToApiHander(&h.JoinModel, p)
		}
	case ApiHanderSqlEasyType_UPDATE:
		{
			hander := &ApiHanderSqlEasySettingUpdate{}
			if h.Setting != "" {
				if err := json.Unmarshal([]byte(h.Setting), hander); err != nil {
					return nil, err
				}
			}
			return hander.ToApiHander(&h.JoinModel, p)
		}
	case ApiHanderSqlEasyType_SEARCHLIST:
		{
			hander := &ApiHanderSqlEasySettingSearchList{}
			if h.Setting != "" {
				if err := json.Unmarshal([]byte(h.Setting), hander); err != nil {
					return nil, err
				}
			}
			return hander.ToApiHander(&h.JoinModel, p)
		}
	case ApiHanderSqlEasyType_SEARCHFIRST:
		{
			hander := &ApiHanderSqlEasySettingSearchFirst{}
			if h.Setting != "" {
				if err := json.Unmarshal([]byte(h.Setting), hander); err != nil {
					return nil, err
				}
			}
			return hander.ToApiHander(&h.JoinModel, p)
		}
	}
	return nil,errors.New("error type")
}