package mysql

import (
	"fmt"
	"miniproject/config"
	"miniproject/model/tables"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = config.Name + ":" + config.PWD + "@tcp" + config.IP + "/miniproject?" + config.Var
var DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

func DataInit() {

	//fmt.Println(dsn)

	if err != nil {
		fmt.Println("DataInit err!", err)
		return
	}

	DB.AutoMigrate(&tables.User{})

	DB.AutoMigrate(&tables.Good{})

	DB.AutoMigrate(&tables.Comment{})

	DB.AutoMigrate(&tables.Cart{})

}
