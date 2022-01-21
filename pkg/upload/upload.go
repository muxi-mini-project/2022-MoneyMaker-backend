package upload

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UploadA(c *gin.Context, goodsid int) string {
	var way string
	name := c.PostForm("name")
	fmt.Println(name)
	file, _, err := c.Request.FormFile("avatar")
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "error!",
		})
		return ""
	}
	//filename := header.Filename
	//fmt.Println(file, err, filename)
	//以goodsid作为名字
	way = "C:\\source\\go\\miniroject\\goods\\avatar" + strconv.Itoa(goodsid) + ".jpg"
	tmp, err := os.Create(way)
	if err != nil {
		fmt.Println(err)
	}
	defer tmp.Close()
	_, err = io.Copy(tmp, file)
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "error!",
		})
		return ""
	}
	fmt.Println("存储成功")
	return way
}

func UploadW(c *gin.Context, goodsid int) string {
	var way string
	name := c.PostForm("name")
	fmt.Println(name)
	file, _, err := c.Request.FormFile("way")
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "error!",
		})
		return ""
	}
	//filename := header.Filename
	//fmt.Println(file, err, filename)
	//以goodsid作为名字
	way = "C:\\source\\go\\miniroject\\goods\\way" + strconv.Itoa(goodsid) + ".jpg"
	tmp, err := os.Create(way)
	if err != nil {
		fmt.Println(err)
	}
	defer tmp.Close()
	_, err = io.Copy(tmp, file)
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "error!",
		})
		return ""
	}
	fmt.Println("存储成功")
	return way
}
