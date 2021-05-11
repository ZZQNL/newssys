package controllers

import (
	//"github.com/beego/beego/v2/client/orm"
	"github.com/astaxie/beego/orm"
	"newspass/models"
	beelog "github.com/beego/beego/v2/logs"
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

type RegController struct{
	beego.Controller
}
type LogController struct{
	beego.Controller
}
func (this*RegController) ShowReg(){
	this.TplName="register.html"
}

func (this*RegController) AddReg(){
	//获取数
	name := this.GetString("userName")
	password := this.GetString("password")
	if name == "" || password == ""{
		beelog.Info("用户或者密码不能为空",name,password)
		this.TplName="register.html"
		return
	}
	o := orm.NewOrm()
	user := models.User{}
	user.Username = name
	user.Password = password
	_,err := o.Insert(&user)
	if err != nil{
		beelog.Info("db inster err",err)
		this.TplName="register.html"
		return
	}
	this.Redirect("/login",302)
	//this.TplName="login.html"
}
func (this*LogController) LogShow(){
	name := this.Ctx.GetCookie("userName")

	if name != ""	{
		this.Data["checked"] = "checked"
		this.Data["userName"] = name
	}
	this.TplName = "login.html"
}

func (this*LogController) Login(){
	name := this.GetString("userName")
	passwd := this.GetString("password")
	check := this.GetString("remember")
	beelog.Info(check)
	if name == "" || passwd == ""{
		beelog.Info("用户密码为空")
		return
	}
	o := orm.NewOrm()
	user := models.User{}
	user.Username = name
	err:=o.Read(&user,"username")
	if err != nil{
		beelog.Info("用户名出错")
		return
	}
	if user.Password != passwd{
		beelog.Info("密码错误")
		return
	}
	if check != "" {
		this.Ctx.SetCookie("userName", name, time.Second*3600)//
	}else {
		this.Ctx.SetCookie("userName",name,-1)//设置有效时间为-1，相当于不保存cookie
	}

	this.SetSession("userName",name)

	this.Redirect("/article/showArticleList",302)
}
func(this * LogController) Logout(){
	this.DelSession("userName")
	this.Redirect("/login",302)
}