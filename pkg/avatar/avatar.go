package avatar

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func GetAvatar() string {
	url := "https://www.duitang.com/blogs/tag/?name=%E6%99%AE%E4%BF%A1%E7%94%B7%E5%A4%B4%E5%83%8F"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(string(body))

	re := `"https://c-ssl.duitang.com/uploads/blog(.*?)"`

	ans := regexp.MustCompile(re)

	str := ans.FindAllString(string(body), -1)

	//fmt.Println(str, len(str), str[1])

	var num int
	rand.Seed(time.Now().UnixNano())
	num = rand.Intn(24)
	for i, v := range str {
		if i == num {
			tmp := strings.Split(v, `"`)
			return tmp[1]
		}
	}
	return ""

}
