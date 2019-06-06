package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/silencily/sparktime/web/controllers"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func newApp() *iris.Application {
	app := iris.Default()

	log := app.Logger()
	log.SetLevel("debug")
	log.SetOutput(os.Stdout)
	log.AddOutput(&lumberjack.Logger{
		Filename:   "./log/spark-time.log",
		MaxSize:    100, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})

	app.RegisterView(iris.HTML("./web/views", ".html"))
	app.StaticWeb("/static", "./web/static")

	mvc.New(app).Handle(new(controllers.IndexController))

	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}
