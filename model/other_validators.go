package model

import (
	"fmt"
	"strings"
)

type Validator map[string][]string

func (v Validator)GenerateGo() string {
	roles := []string{}
	for k, t := range v {
		for i := 0; i < len(t); i++ {
			t[i] = fmt.Sprintf("\"%s\"", t[i])
		}
		role := fmt.Sprintf("\t\t\"%s\":{%s},", k, strings.Join(t, ","))
		roles = append(roles, role)
	}
	if len(roles) <= 0 {
		return ""
	}
	str := fmt.Sprintf("\tvar roleValidators = map[string][]string{\n%s\n\t}\n", strings.Join(roles, "\n"))
	str += "\tif err := utils.Verify(*reqStruct,roleValidators);err != nil{\n\t\thttpx.FailWithMessage(c,err.Error())\n\t\treturn\n\t}\n"
	return str
}