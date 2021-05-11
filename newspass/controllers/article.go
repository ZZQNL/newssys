package controllers

import (
	//"github.com/beego/beego/v2/client/orm"
	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
	beelog "github.com/astaxie/beego/logs"
	"math"
	"newspass/models"
	"path"
	"time"
)

type ArticlCon struct{
	beego.Controller
}
//文章展示列表
//下拉框的功能业务
func (this*ArticlCon) ArtTSelect(){
	sel := this.GetString("select")
	//查询数据
	o := orm.NewOrm()
	var article []models.Article
	//数据表关联查询,添加过滤
	if sel == ""{
		o.QueryTable("Article").RelatedSel("ArticleType").All(&article)
	}else{
		o.QueryTable("Article").RelatedSel("ArtType").Filter("ArtType__Typename",sel).All(&article)
	}
	this.Data["typeName"] = sel
	this.Redirect("/article/index",302)
}

func (this*ArticlCon) Artshow(){
	o := orm.NewOrm()
	var articlewithtype[]models.Article
	sel := this.GetString("select")
	res := o.QueryTable("Article")
	//res.All(&article)

	//页码逻辑
	//pageIndex := 1//默认为1为首页

	pageIndex,err:= this.GetInt("pageIndex")//更新后的页码
	if err != nil{
		pageIndex = 1//如果没有传值，就访问第一页
	}
	//判断select值是否为空
	var count int64
	if sel !=""{
		count,_ = res.RelatedSel("ArtType").Filter("ArtType__Typename",sel).Count()

	}else {
		count,_ = res.Count()
	}

	//count,_ := res.RelatedSel("ArtType").Count()

	pageSize := 2
	start := pageSize*(pageIndex-1)
	res.Limit(pageSize,start).RelatedSel("ArtType").All(&articlewithtype)
	pageCount := float64(count)/float64(pageSize)
	pageCount1 := math.Ceil(pageCount)
	pI := float64(pageIndex)
	//页码超出范围逻辑
	//上一页超出范围
	pageFirst := false
	if pI <= 1{
		pageFirst = true
		pageIndex = 1
	}
	this.Data["pageFirst"] = pageFirst
	//下一页超出范围
	pageEnd := false
	if pI >= pageCount1{
		pageEnd = true
		pageIndex = int(pageCount1)
	}

	//var articlewithtype []models.Article

	/*------------------------根据选中的类型查询响应的文章----------------------------------------*/
	if sel != ""{
		res.Limit(pageSize,start).RelatedSel("ArtType").Filter("ArtType__Typename",sel).All(&articlewithtype)
	}else{
		res.Limit(pageSize,start).RelatedSel("ArtType").All(&articlewithtype)
	}

	var at []models.ArtType
	o.QueryTable("ArtType").All(&at)

	this.Data["articleTypes"] = at
	this.Data["typeName"] = sel
	this.Data["pageEnd"] = pageEnd
	this.Data["articles"] = articlewithtype
	//beelog.Info(article)
	this.Data["pageCounts"] = pageCount1
	this.Data["pageIndex"] = pageIndex
	this.Data["count"] = count
	this.Data["userName"] = this.GetSession("userName")
	this.Layout = "layout.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["conthead"] = "head.html"
	this.TplName = "index.html"
}
func (this*ArticlCon) ArtAddShow(){
	o := orm.NewOrm()
	var at []models.ArtType
	_,err := o.QueryTable("ArtType").All(&at)
	if err != nil{
		beelog.Info("查询错误")
		return
	}
	this.Data["articleType"] = at
	this.Layout = "layout.html"
	this.TplName="add.html"
}

//插入文章
func (this*ArticlCon) ArtAdd(){
	f,h,err :=this.GetFile("uploadname")
	//beelog.Info("head:",h,"file:",f)
	defer f.Close()
	t := time.Now().Format("2006-01-02 15:04:05")
	if err != nil{
		beelog.Info("读取文件出错")
		return
	}
	ext := path.Ext(h.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg"{
		beelog.Info("文件的格式不正确")
		this.TplName = "add.html"
		return
	}
	this.SaveToFile("uploadname","./static/img/"+t+ext)
	//beelog.Info("保存成功",h.Filename)
	typename := this.GetString("select")//下拉框的值
	artname := this.GetString("articleName")
	content := this.GetString("content")
	img := "./static/img/"+t+ext
	o := orm.NewOrm()
	article := models.Article{}

	at := models.ArtType{}
	at.Typename = typename
	o.Read(&at,"Typename")//对应字段查询匹配的行
	//beelog.Info(at)
	//beelog.Info(at.Typename)
	article.ArtType = &at
	//beelog.Info(article.ArtType.Typename)
	//article.Time = t
	article.Content = content
	article.Img = img
	article.Title = artname
	_,err2 := o.Insert(&article)
	if err2 != nil {
		beelog.Info("插入出错:",err2)
		this.TplName = "add.html"
		return
	}
	//beelog.Info(article)
	//beelog.Info(article.ArtType)
	//beelog.Info(article.ArtType.Typename)
	this.Redirect("index",302)
	//this.TplName = "add.html"
}

//删除
func (this*ArticlCon) DelArt(){
	id,_ := this.GetInt("id")
	o := orm.NewOrm()
	art := models.Article{Id:id}
	o.Delete(&art)
	this.Redirect("/article/index",302)

}


//更新
func (this*ArticlCon) ShowUpdate(){
	id,_ := this.GetInt("id")
	o := orm.NewOrm()
	art := models.Article{Id:id}
	err := o.Read(&art)
	if err != nil {
		beelog.Info(err)
		this.TplName = "index.html"
		return
	}
	beelog.Info(art.Img)
	this.Data["article"] = art
	this.Layout = "layout.html"
	this.TplName = "update.html"
}
func (this*ArticlCon) PosUpdate(){
	id,_ := this.GetInt("id")
	cont := this.GetString("content")
	title := this.GetString("articleName")
	f,h,_ := this.GetFile("uploadname")
	defer f.Close()

	ext := path.Ext(h.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg"{
		beelog.Info("文件的格式不正确")
		this.TplName = "update.html"
		return
	}
	this.SaveToFile("uploadname","./static/img/"+h.Filename)
	beelog.Info("保存成功",h.Filename)
	img := "./static/img/"+h.Filename
	o := orm.NewOrm()
	art := models.Article{}
	art.Title = title
	art.Content = cont
	art.Id = id
	art.Img = img
	_,err :=o.Update(&art)
	if err != nil {
		beelog.Info("更新出错！")
		return
	}
	this.Redirect("/article/index",302)
}


//文章类型
type ArtTypeCon struct {
	beego.Controller
}
func (this *ArtTypeCon) ShowType(){
	o := orm.NewOrm()
	var aType []models.ArtType
	_,err := o.QueryTable("ArtType").All(&aType)
	if err != nil {
		beelog.Info("没有记录")
		return
	}
	beelog.Info(aType)
	this.Data["data"] = aType
	this.Layout = "layout.html"
	this.TplName = "addType.html"
}
func (this *ArtTypeCon) AddArtType(){
	tname := this.GetString("typeName")
	if tname == "" {
		this.Redirect("/article/addType",302)
		return
	}
	o := orm.NewOrm()
	at := models.ArtType{}
	at.Typename = tname
	o.Insert(&at)
	this.Redirect("/article/addType",302)
}
