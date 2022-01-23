package easy

import (
	"miniproject/model/mysql"
	"miniproject/model/tables"
)

func Returnstar(buyer string, owner string) {
	var msg tables.Message
	msg.Buyer = buyer
	msg.My = owner
	msg.Msg = buyer + "收藏了你的商品"
	mysql.DB.Create(&msg)
}

func Returnbuy(buyer string, owner string) {
	var msg tables.Message
	msg.Buyer = buyer
	msg.My = owner
	msg.Msg = buyer + "购买了你的商品"
	mysql.DB.Create(&msg)
}
