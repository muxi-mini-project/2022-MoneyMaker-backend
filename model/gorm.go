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

func GetOrderCart(id string) (tables.Cart, error) {
	var cart tables.Cart
	err := mysql.DB.Where("id=?", id).Find(&cart).Error
	return cart, err
}

func GetOrderGood(goodsidint int) (tables.Good, error) {
	var good tables.Good
	err := mysql.DB.Where("goods_id=?", goodsidint).Find(&good).Error
	return good, err
}

func GetOrderUser(id string) (tables.User, error) {
	var user tables.User
	err := mysql.DB.Model(&tables.User{}).Where("id=?", id).Find(&user).Error
	return user, err
}

func UpdateBuygoods(id string, buygoods string) error {
	return mysql.DB.Model(&tables.User{}).Where("id=?", id).Update("buygoods", buygoods).Error
}

func UpdateGoodBuyer(id int, buyer string) error {
	return mysql.DB.Model(&tables.Good{}).Where("goods_id=?", id).Update("buyer", buyer).Error
}

func GetGoodComment(id int) ([]tables.Comment, error) {
	var comments []tables.Comment
	err := mysql.DB.Model(&tables.Comment{}).Where("goods_id=?", id).Find(&comments).Error
	return comments, err
}
