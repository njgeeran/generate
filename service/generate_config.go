package service

import (
	"fmt"
	"github.com/njgeeran/generate/model"
	"os"
	"text/template"
)

//生成配置文件
var config_temp = `#配置文件
ENV: local
HTTPADDR: 8080
HOST: 127.0.0.1

{{.setting}}
`

func generate_config(msg chan string,root_path string,p *model.Project) error {
	config_file_path := root_path+"/config.yaml"
	SendGenerateMsg(msg,"开始生成配置文件:"+config_file_path)

	//数据库配置
	setting := "#数据库\nDB: \n  "
	ds := map[model.DataSourceType][]model.DataSource{}
	for _,t := range p.DataSources {
		if ds[t.Type] == nil || len(ds[t.Type]) <= 0 {
			ds[t.Type] = []model.DataSource{t}
		}else {
			ds[t.Type] = append(ds[t.Type], t)
		}
	}
	for k,v := range ds {
		switch k {
		case model.DataSourceType_MsSql:
			setting+= "MSSQL: \n    "
			break
		case model.DataSourceType_MySql:
			setting+= "MYSQL: \n    "
			break
		}
		for _,t := range v {
			setting+= fmt.Sprintf(`%s: 
      username: %s
      password: %s
      path: %s
      db_name: %s
      config: %s
      max-idle-conns: %d
      max-open-conns: %d`,t.Name,t.UserName,t.Password,t.Path,t.DbName,t.Config,t.MaxIdleConns,t.MaxOpenConns)
		}
	}

	//创建配置文件
	f,err := os.Create(config_file_path)
	if err != nil {
		SendGenerateErrMsg(msg,"配置文件创建失败:",err)
		return err
	}
	tmpl := template.Must(template.New("config").Parse(config_temp))
	tmpl.Execute(f, map[string]string{
		"setting":setting,
	})
	f.Close()
	return nil
}