package model

import (
	"miniproject/model/mysql"
	"miniproject/model/tables"
)

func GetLastRecord() (tables.Good, error) {
	var good tables.Good
	err := mysql.DB.Model(&tables.Good{}).Order("goods_id desc").Take(&good).Error
	return good, err
}

func GetOrderCart(id string) tables.Cart {
	var cart tables.Cart
	mysql.DB.Where("id=?", id).Find(&cart)
	return cart
}

func GetOrderGood(goodsidint int) tables.Good {
	var good tables.Good
	mysql.DB.Where("goods_id=?", goodsidint).Find(&good)
	return good
}

func GetOrderUser(id string) tables.User {
	var user tables.User
	mysql.DB.Model(&tables.User{}).Where("id=?", id).Find(&user)
	return user
}

func UpdateBuygoods(id string, buygoods string) error {
	return mysql.DB.Model(&tables.User{}).Where("id=?", id).Update("buygoods", buygoods).Error
}

func UpdateGoodBuyer(id int, buyer string) error {
	return mysql.DB.Model(&tables.Good{}).Where("goods_id=?", id).Update("buyer", buyer).Error
}

func GetGoodComment(id int) []tables.Comment {
	var comments []tables.Comment
	mysql.DB.Model(&tables.Comment{}).Where("goods_id=?", id).Find(&comments)
	return comments
}
