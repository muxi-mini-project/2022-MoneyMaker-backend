package upload

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UploadA(c *gin.Context, goodsid int) bool {
	var way string
	name := c.PostForm("name")
	fmt.Println(name)
	file, _, err := c.Request.FormFile("avatar")
	if err != nil {
		fmt.Println(err)
		return false
	}
	//filename := header.Filename
	//fmt.Println(file, err, filename)
	//以goodsid作为名字
	way = "C:\\source\\go\\miniroject\\goods\\avatar" + strconv.Itoa(goodsid) + ".jpg"
	tmp, err := os.Create(way)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer tmp.Close()
	_, err = io.Copy(tmp, file)
	if err != nil {
		fmt.Println(err)
		return false
	}
	//fmt.Println("存储成功")
	return true
}

func UploadW(c *gin.Context, goodsid int) bool {
	var way string
	name := c.PostForm("name")
	fmt.Println(name)
	file, _, err := c.Request.FormFile("way")
	if err != nil {
		fmt.Println(err)
		return false
	}
	//filename := header.Filename
	//fmt.Println(file, err, filename)
	//以goodsid作为名字
	way = "C:\\source\\go\\miniroject\\goods\\way" + strconv.Itoa(goodsid) + ".jpg"
	tmp, err := os.Create(way)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer tmp.Close()
	_, err = io.Copy(tmp, file)
	if err != nil {
		fmt.Println(err)
		return false
	}
	//fmt.Println("存储成功")
	return true
}
