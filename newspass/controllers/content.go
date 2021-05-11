package controllers

import (
	//"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
	"newspass/models"
	"strconv"
	beelog "github.com/beego/beego/v2/logs"
)

type ContCtr struct{
	beego.Controller
}
func (this*ContCtr) ContShow(){
	id := this.GetString("id")
	o := orm.NewOrm()
	id2,_ := strconv.Atoi(id)
	article := models.Article{Id:id2}
	err := o.Read(&article)
	if err != nil {
		beelog.Info("数不存在")
		return
	}
	article.Count += 1
	o.Update(&article)

	/* 多对多插入*/
	//获取多对多的对象
	m2m := o.QueryM2M(&article,"Users")
	userName := this.GetSession("userName")
	user := models.User{}
	user.Username = userName.(string)
	o.Read(&user,"UserName")

	_,err = m2m.Add(&user)
	if err != nil {
		beelog.Info("插入失败")
		return
	}
	o.Update(&article)//没有指定更新哪一列也没事，会自动寻找

	var users []models.User
	o.QueryTable("User").Filter("Articles__Article__Id",id2).Distinct().All(&users)

	this.Data["users"] = users
	//beelog.Info(article.ArtType)
	this.Data["article"]=article
	this.Layout = "layout.html"
	this.TplName = "content.html"

}