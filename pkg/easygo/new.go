package easy

import (
	"strings"
)

func New(str string, id string) string {
	var re string
	strstr := strings.Split(str, ",")
	//fmt.Println(strstr)
	new := append(strstr, id)
	//这里new是append之后的，长度已经加一了
	for i, v := range new {
		if i < len(new)-1 {
			re = re + v + ","
		} else {
			re = re + v
		}
		//fmt.Println(re)
	}
	//fmt.Println(re)
	return re
}

func NewSingle(str string, id string) (string, bool) {
	var re string
	strstr := strings.Split(str, ",")
	//fmt.Println(strstr)
	new := append(strstr, id)
	//这里new是append之后的，长度已经加一了
	for i, v := range new {
		if i < len(new)-1 {
			if v == id {
				return str, false
			}
			re = re + v + ","
		} else {
			re = re + v
		}
		//fmt.Println(re)
	}
	//fmt.Println(re)
	return re, true
}
