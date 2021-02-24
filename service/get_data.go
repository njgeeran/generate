package service

import (
	"github.com/jinzhu/gorm"
	"lowcode/generate/model"
)

func GetProjectData(db *gorm.DB,pid int) (*model.Project,error) {
	var err error

	p := &model.Project{}
	if err := db.First(p,"id = ?",pid).Error;err != nil{
		return nil,err
	}
	//获取model
	if p.Models,err = get_model_data(db,p);err != nil{
		return nil,err
	}
	//获取数据源
	if p.DataSources,err = get_datasource_data(db,p);err != nil{
		return nil,err
	}
	//获取引用模块
	if p.Modules,err = get_module_data(db,p);err != nil{
		return nil,err
	}
	//获取路由
	if p.Routers,err = get_router_data(db,p);err != nil {
		return nil,err
	}
	return p,nil
}

func get_model_data(db *gorm.DB,p *model.Project) ([]model.Model,error) {
	ms := []model.Model{}
	var err error

	if err := db.Find(&ms,"from_project = ?",p.ID).Error;err != nil{
		return nil,err
	}

	//获取字段
	for i:=0 ;i< len(ms);i++  {
		if ms[i].Fields,err = get_field_data(db,&ms[i]);err != nil{
			return nil,err
		}
	}
	return ms,nil
}
func get_field_data(db *gorm.DB,m *model.Model) ([]model.Field,error) {
	fs := []model.Field{}
	if err := db.Find(&fs,"from_model = ?",m.ID).Error;err != nil{
		return nil,err
	}
	return fs,nil
}

func get_datasource_data(db *gorm.DB,p *model.Project) ([]model.DataSource,error) {
	dss := []model.DataSource{}

	if err := db.Find(&dss,"from_project = ?",p.ID).Error;err != nil{
		return nil,err
	}

	for i:=0 ; i< len(dss) ; i++ {
		for _,t := range p.Models {
			if t.FromDS == dss[i].ID {
				if dss[i].Models == nil {
					dss[i].Models = []model.Model{t}
				}else {
					dss[i].Models = append(dss[i].Models, t)
				}
			}
		}
	}
	return dss,nil
}

func get_module_data(db *gorm.DB,p *model.Project) ([]model.Module,error) {
	ms := []model.Module{}
	if err := db.Find(&ms,"from_project = ? or from_project = ''",p.ID).Error;err != nil{
		return nil,err
	}
	return ms,nil
}

func get_router_data(db *gorm.DB,p *model.Project) ([]model.Router,error) {
	rs := []model.Router{}
	if err := db.Find(&rs,"from_project = ?",p.ID).Error;err != nil{
		return nil,err
	}

	var err error
	for i:=0 ; i<len(rs) ; i++  {
		switch rs[i].HanderType {
		case model.ApiHanderType_DIY:
			if rs[i].Hander,err = get_api_hander_diy_data(db,p,&rs[i]);err != nil{
				return nil,err
			}
		case model.ApiHanderType_EasySql:
			if rs[i].Hander,err = get_api_hander_sql_easy_data(db,p,&rs[i]);err != nil{
				return nil,err
			}
		}
	}
	return rs,nil
}
func get_api_hander_diy_data(db *gorm.DB,p *model.Project,r *model.Router) (*model.ApiHander,error) {
	diy := model.ApiHanderDiy{}
	if err := db.First(&diy,"id = ?",r.HanderId).Error;err != nil{
		return nil,err
	}
	return diy.ToApiHander(p)
}
func get_api_hander_sql_easy_data(db *gorm.DB,p *model.Project,r *model.Router) (*model.ApiHander,error) {
	sql_easy := model.ApiHanderSqlEasy{}
	if err := db.First(&sql_easy,"id = ?",r.HanderId).Error;err != nil{
		return nil,err
	}
	return sql_easy.ToApiHander(p)
}