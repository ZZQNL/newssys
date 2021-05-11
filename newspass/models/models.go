package models

import (
	//"github.com/beego/beego/v2/client/orm"
	"time"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id int
	Username string
	Password string
	Articles []*Article `orm:"reverse(many)"`//
}
type Article struct {
	Id      int       `orm:"pk;auto"`
	Title   string    `orm:"size(20)"`
	Kind    string    `orm:"size(20)"`
	Content string    `orm:"size(500)"`
	Img     string    `orm:"size(50);null"`
	Time    time.Time `orm:"type(datatime);auto_now_add"` //自动添加当前时间
	Count   int       `orm:"default(0)"`
	ArtType *ArtType	`orm:"rel(fk)"`//外键
	Users []*User	`orm:"rel(m2m)"`//
}
type ArtType struct {
	Id int
	Typename string		`orm:"size(20)"`
	Article []*Article	`orm:"reverse(many)"`//一对多的反向关系
}
//func init(){
//	//连接数据库
//	orm.Debug=true
//	orm.RegisterDataBase("default","mysql","root:123456@tcp(localhost:3306)/userpass?charset=utf8")
//	//注册表
//	orm.RegisterModel(new(User))
//	//orm.RegisterModel(new)
//	//生成表
//	orm.RunSyncdb("default",false,true)
//}
func init(){
	//连接数据库
	orm.Debug=true
	orm.RegisterDataBase("default","mysql","root:123456@tcp(localhost:3306)/userpass?charset=utf8")
	//注册表
	//orm.RegisterModel(new(User))
	orm.RegisterModel(new(Article),new(User),new(ArtType))
	//生成表
	orm.RunSyncdb("default",false,true)
}