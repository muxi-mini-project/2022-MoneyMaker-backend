package mysql

import (
	"fmt"
	"miniproject/config"
	"miniproject/model/tables"
	"miniproject/pkg/avatar"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//防止第一个新增商品出现错误
func DataInit() {

	DB.Model(tables.User{}).Create(map[string]interface{}{
		"id": "0",
	})
	fmt.Println("user已初始化")

	DB.Model(&tables.Good{}).Create(map[string]interface{}{
		"goods_id": 0,
		//"price":    0,
	})
	fmt.Println("goods已初始化")

	DB.Model(&tables.Message{}).Create(map[string]interface{}{
		"id": 0,
	})
	fmt.Println("message已初始化")

}

var dsn = config.Name + ":" + config.PWD + "@tcp" + config.IP + "/miniproject?" + config.Var
var DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

func Init() {

	//fmt.Println(dsn)

	//新建一个空的通知，id的大小可以分辨出其时间先后

	if err != nil {
		fmt.Println("Init err!", err)
		return
	}

	DB.AutoMigrate(&tables.User{})

	DB.AutoMigrate(&tables.Good{})

	DB.AutoMigrate(&tables.Comment{})

	DB.AutoMigrate(&tables.Cart{})

	DB.AutoMigrate(&tables.Message{})

	DataInit()

}

//登录成功初始化这个用户的信息
func Create(id string, name string) {
	var user tables.User
	//不存在就会报错
	DB.Where("id=?", id).Find(&user)
	avatar := avatar.GetAvatar()
	if user.ID != id {
		DB.Model(&tables.User{}).Create(map[string]interface{}{
			"id":       id,
			"avatar":   avatar, //随机分配头像
			"nickname": name,
			"buygoods": "",
		})

		//新建一个空的购物车
		DB.Model(&tables.Cart{}).Create(map[string]interface{}{
			"id": id})
	}
}
