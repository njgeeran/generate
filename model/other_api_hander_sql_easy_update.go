package model

import (
	"errors"
	"fmt"
	"github.com/njgeeran/core/utils"
	"strconv"
)

type ApiHanderSqlEasySettingUpdate struct {
	Validators 	Validator `json:"validators"`
}

func (h *ApiHanderSqlEasySettingUpdate)ToApiHander(m *Model,p *Project) (*ApiHander,error) {
	ah := &ApiHander{}
	ah.BodyParaModel = &BodyPara{
		ModelId:m.ID,
		JoinModel:*m,
		Validators:h.Validators,
		ContentType:BodyParaContentType_JSON,
	}
	ah.BodyParaModel.Validators["ID"] = []string{utils.NotEmpty()}
	ah.ReturnVal = nil
	ds := &DataSource{}
	for _,t := range p.DataSources {
		if t.ID == m.FromDS {
			ds = &t
			break
		}
	}
	if ds == nil {
		return nil,errors.New("can't find datasource by model["+strconv.Itoa(m.ID)+"]")
	}
	ah.HanderCode = fmt.Sprintf("\terr = model.GetDb%s().Save(reqStruct).Error\n\treturn", utils.UnderlinToCamelCase(ds.Name))
	return ah,nil
}
