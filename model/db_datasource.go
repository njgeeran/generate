package model

import "github.com/njgeeran/core/utils"

type DataSource struct {
	utils.BaseModel
	FromProject int `json:"from_project"` //属于项目

	Type DataSourceType `json:"type"` //数据源类型

	Name string `json:"name"` //名称【英文,首字母大写】

	UserName     string `json:"user_name"`
	Password     string `json:"password"`
	Path         string `json:"path"`
	DbName       string `json:"db_name"`
	Config       string `json:"config"`
	MaxIdleConns int    `json:"max_idle_conns"`
	MaxOpenConns int    `json:"max_open_conns"`

	SingularTable bool `json:"singular_table"` //禁用表名复数

	Models []Model
}
type DataSourceType int
const (
	_DataSourceType = iota
	DataSourceType_MySql
	DataSourceType_MsSql
)
