package service

import (
	"github.com/jinzhu/gorm"
	"github.com/njgeeran/core/conf"
	"github.com/njgeeran/core/orm"
	"os"
)

func InitTest()  {
	os.Chdir("../")
	conf.InitConfg()
	orm.InitOrm(conf.GetConf())
	db := orm.GetOrm().Get("low_code")
	db.SingularTable(true)
}
func GetOrm() *gorm.DB {
	return orm.GetOrm().Get("low_code")
}
