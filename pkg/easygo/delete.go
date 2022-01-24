package easy

import (
	"strings"
)

func Delete(str string, id string) string {
	var re string
	var count = 0

	strstr := strings.Split(str, ",")
	//fmt.Println(strstr, len(strstr))

	for i, v := range strstr {
		var ok = false
		if v == id {
			ok = true
			count++
			if ok && count == 1 {
				continue
			}
		}

		if i < len(strstr)-1 {
			re = re + v + ","
		} else {
			re = re + v
		}

		//fmt.Println(re)
	}
	//fmt.Println(re)
	return re
}
