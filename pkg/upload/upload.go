package upload

import (
	"io"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UploadAvatar(c *gin.Context, goodsid int) bool {
	var way string

	file, _, err := c.Request.FormFile("avatar")
	if err != nil {
		log.Println(err)
		return false
	}
	//filename := header.Filename
	//fmt.Println(file, err, filename)
	//以goodsid作为名字
	way = "./goods\\avatar\\" + strconv.Itoa(goodsid) + ".jpg"
	tmp, err := os.Create(way) //如果文件已存在会清空

	if err != nil {
		log.Println(err)
		return false
	}

	defer tmp.Close()

	_, err = io.Copy(tmp, file)
	if err != nil {
		log.Println(err)
		return false
	}
	//fmt.Println("存储成功")
	return true
}

func UploadWay(c *gin.Context, goodsid int) bool {
	var way string

	file, _, err := c.Request.FormFile("way")

	if err != nil {
		log.Println(err)
		return false
	}
	//filename := header.Filename
	//fmt.Println(file, err, filename)
	//以goodsid作为名字
	way = "./goods\\way\\" + strconv.Itoa(goodsid) + ".jpg"

	tmp, err := os.Create(way)
	if err != nil {
		log.Println(err)
		return false
	}

	defer tmp.Close()

	_, err = io.Copy(tmp, file)
	if err != nil {
		log.Println(err)
		return false
	}
	//fmt.Println("存储成功")
	return true
}
