package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	_ "newspass/routers"
)

func main() {
	beego.AddFuncMap("prePage",showPre)//上一页映射
	beego.AddFuncMap("nextPage",showNext)//下一页映射
	beego.Run()
}

//上一页，视图函数
func showPre(data int) int{
	data = data-1
	return data
}
func showNext(data int) int{
	data = data+1
	return data
}