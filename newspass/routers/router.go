package routers

import (
	//"github.com/astaxie/beego/context"
	"newspass/controllers"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func init() {
	//在路由之前先过滤
	beego.InsertFilter("/article/*",beego.BeforeRouter,FilterFunc)
    beego.Router("/", &controllers.MainController{})
	beego.Router("/register",&controllers.RegController{},"get:ShowReg;post:AddReg")
	beego.Router("/login",&controllers.LogController{},"get:LogShow")
    beego.Router("/login",&controllers.LogController{},"post:Login")
    beego.Router("/article/index",&controllers.ArticlCon{},"get:Artshow;post:ArtTSelect")
    beego.Router("/article/add",&controllers.ArticlCon{},"post:ArtAdd")
	beego.Router("/article/add",&controllers.ArticlCon{},"get:ArtAddShow")
    beego.Router("/article/showArticleDetail",&controllers.ContCtr{},"get:ContShow")
    beego.Router("/article/delArticle",&controllers.ArticlCon{},"get:DelArt")
	beego.Router("/article/showUpdate",&controllers.ArticlCon{},"get:ShowUpdate;post:PosUpdate")
    beego.Router("/article/showArticleList",&controllers.ArticlCon{},"get:Artshow;post:ArtTSelect")
    beego.Router("/article/addType",&controllers.ArtTypeCon{},"get:ShowType;post:AddArtType")
	beego.Router("/logout",&controllers.LogController{},"get:Logout")
    //beego.Router("addType",&controllers.ArticlCon{},"get:showType")
    //beego.Router("/article/showArtType",&controllers.ArticlCon{},"post:showArtType")

}
var FilterFunc = func(ctx *context.Context){
	name := ctx.Input.Session("userName")
	if name == nil{
		ctx.Redirect(302,"/login")
	}
}
