package tables

//"gorm.io/driver/mysql"

type User struct {
	//gorm.Model
	Nickname string
	ID       string
	Password string
	Avatar   string
}

//下架了，但是用户如果已经接单完成后，该用户任然可以查看到，只是无法被搜索以及从我的橱窗显示
type Good struct {
	GoodsID  int ``
	ID       string
	Scores   int //总评分
	Summary  string
	Zone     string //分区
	Price    int
	in       string
	FeedBack int    //被举报的次数
	Way      string //联系方式，但感觉一个商品一个联系方式好奇怪啊,存放图片
	Buyer    string //一个用户只能购买一次，但可以被多个用户购买
}

type Comment struct {
	//gorm.Model
	//CommentID int `gorm:"primary key"`,感觉没啥必要了
	Comment string
	Score   int //一个用户的单独评分
	GoodsID int
	ID      string
}

type Cart struct {
	ID    string
	Buyer string
}
