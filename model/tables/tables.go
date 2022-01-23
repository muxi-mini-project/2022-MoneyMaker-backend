package tables

//"gorm.io/driver/mysql"

type User struct {
	Nickname string `gorm:"nickname"`
	ID       string `gorm:"id"`
	Avatar   string `gorm:"avatar"`
	Buygoods string `gorm:"buygoods"`
}

//下架了，就都差看不到了
type Good struct {
	GoodsID   int     `gorm:"goods_id"`
	Title     string  `gorm:"title"`
	ID        string  `gorm:"id"`
	Summary   string  `gorm:"summary"`
	Goodszone string  `gorm:"goodszone"`
	Price     int     `gorm:"price"`
	Scores    float64 `gorm:"scores"`
	Goodsin   string  `gorm:"goodsin"`
	Avatar    string  `gorm:"avatar"`    //随机分配头像，直接在网上找几个就可以了，所以只要存网址
	FeedBack  int     `gorm:"feed_back"` //被举报的次数
	Way       string  `gorm:"way"`       //联系方式，但感觉一个商品一个联系方式好奇怪啊,存放图片
	Buyer     string  `gorm:"buyer"`     //用户确认完成后则会把他的名字删除
}

type Comment struct {
	CommentID int    `gorm:"comment_id"`
	Comment   string `gorm:"comment"`
	Score     int    `gorm:"score"`
	GoodsID   int    `gorm:"goods_id"`
	ID        string `gorm:"id"`
	Givetime  string `gorm:"givetime"`
}

type Cart struct {
	ID      string `gorm:"id"`
	Goodsid string `gorm:"goodsid"`
}

type Message struct {
	Buyer string `gorm:"buyer"`
	My    string `gorm:"my"`
	Msg   string `gorm:"msg"`
	ID    int    `gorm:"id"`
}
