package main

import (
	"context"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"go_Iris/backend/web/controllers"
	"go_Iris/common"
	"go_Iris/repositories"
	"go_Iris/services"
)

func main() {
	// 1.创建Iris实例
	app := iris.New()

	// 2.设置错误模式，在mvc模式下提示错误
	app.Logger().SetLevel("debug")

	// 3.注册模板
	// 指定解析模板的目录，为html后缀 布局模板目录
	tmplate := iris.HTML("./backend/web/views", ".html").Layout("shared/layout.html").Reload(true)
	//app.RegisterView(iris.HTML("./backend/web/views", ".html"))

	app.RegisterView(tmplate)
	// 4.注册静态资源
	app.HandleDir("/assets", iris.Dir("./backend/web/assets"))

	// 5.出现异常跳转到指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问的页面出错！"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})
	//连接数据库
	db, err := common.NewMysqlConn()
	if err != nil {
		panic(err.Error())
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//5.注册控制器
	productRepository := repositories.NewProductManager("product", db)
	productSerivce := services.NewProductService(productRepository)
	productParty := app.Party("/product")
	product := mvc.New(productParty)
	product.Register(ctx, productSerivce)
	product.Handle(new(controllers.ProductController))

	//6.启动服务
	app.Run(
		iris.Addr("localhost:8080"),
	)

}
