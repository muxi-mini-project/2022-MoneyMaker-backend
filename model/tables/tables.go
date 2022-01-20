package tables

//"gorm.io/driver/mysql"

type User struct {
	Nickname string `gorm:"nickname"`
	ID       string `gorm:"id"`
	Avatar   string `gorm:"avatar"`
}

//下架了，就都差看不到了
type Good struct {
	GoodsID   int    `gorm:"goodsid"`
	ID        string `gorm:"id"`
	Scores    int    `gorm:"scores"`
	Summary   string `gorm:"summary"`
	Goodszone string `gorm:"goodszone"`
	Price     int    `gorm:"price"`
	Goodsin   string `gorm:"goodsin"`
	Avatar    string
	FeedBack  int    `gorm:"feedback"` //被举报的次数
	Way       string `gorm:"way"`      //联系方式，但感觉一个商品一个联系方式好奇怪啊,存放图片
	Buyer     string `gorm:"buyer"`    //用户确认完成后则会把他的名字删除
}

type Comment struct {
	CommentID int    `gorm:"commentid"`
	Comment   string `json:"comment" binding:"required"`
	Score     int    `json:score binding:"required"`
	GoodsID   int    `gorm:"goodsid"`
	ID        string `gorm:"id"`
}

type Cart struct {
	ID      string `gorm:"id"`
	Goodsid string `gorm:"goodsid"`
}
