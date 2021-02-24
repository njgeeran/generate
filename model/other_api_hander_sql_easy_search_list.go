package model

import (
	"errors"
	"fmt"
	"github.com/njgeeran/core/utils"
	"strconv"
)

type ApiHanderSqlEasySettingSearchList struct {
	UrlParas 		[]struct{
		UrlPara
		ModelFieldName 		string			`json:"model_field_name"`
	}			`json:"url_paras"`
}
func (h *ApiHanderSqlEasySettingSearchList)ToApiHander(m *Model,p *Project) (*ApiHander,error) {
	ah := &ApiHander{}
	ah.BodyParaModel = nil
	ah.ReturnVal = &ReturnVal{
		FieldType:"httpx.BasePageModel",
	}
	ah.UrlParaModel = append(ah.UrlParaModel, UrlPara{
		FieldName:"limit",FieldType:"int",DefaultValue:"10",
	}, UrlPara{
		FieldName:"page",FieldType:"int",DefaultValue:"0",
	})

	//获取数据源
	ds := &DataSource{}
	{
		for _,t := range p.DataSources {
			if t.ID == m.FromDS {
				ds = &t
				break
			}
		}
		if ds == nil {
			return nil,errors.New("can't find datasource by model["+strconv.Itoa(m.ID)+"]")
		}
	}

	hander_str := ""
	{
		hander_str = fmt.Sprintf("\tdb := model.GetDb%s()\n",utils.UnderlinToCamelCase(ds.Name))
		hander_str += fmt.Sprintf("\tdb = db.Model(&model.%s{})\n",m.Name)
		hander_str += fmt.Sprintf("\tdata = &httpx.BasePageModel{}\n\td := &[]model.%s{}\n\tdata.Data = d\n",m.Name)
		for _,t := range h.UrlParas {
			ah.UrlParaModel = append(ah.UrlParaModel, t.UrlPara)
			switch t.FieldType {
			case "int":
				hander_str += fmt.Sprintf("\tif %s != 0 {\n\tdb = db.Where(\"%s = ?\",%s)\n\t}\n", t.FieldName, t.ModelFieldName, t.FieldName)
			case "string":
				hander_str += fmt.Sprintf("\tif %s != \"\" {\n\tdb = db.Where(\"%s = ?\",%s)\n\t}\n", t.FieldName, t.ModelFieldName, t.FieldName)
			}
		}
		hander_str += "\tdb.Count(&data.Count)\n"
		return_str := fmt.Sprintf("\terr = db.Limit(limit).Offset((page - 1) * limit).Find(d).Error\n\treturn")
		hander_str += return_str
	}

	ah.HanderCode = hander_str
	return ah,nil
}