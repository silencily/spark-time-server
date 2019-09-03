package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"github.com/silencily/sparktime/core"
	"github.com/silencily/sparktime/core/task"
	"github.com/silencily/sparktime/services"
	"github.com/silencily/sparktime/web/controllers"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

const maxSize = 7 << 19 // 600KB

func newApp() *iris.Application {
	app := iris.Default()

	log := app.Logger()
	log.SetLevel("debug")
	log.SetOutput(os.Stdout)
	log.SetTimeFormat("2006/01/02 15:04:05")
	log.AddOutput(&lumberjack.Logger{
		Filename:   "./log/spark-time.log",
		MaxSize:    100, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})
	core.SetRootLogger(log)

	app.RegisterView(iris.HTML("./web/views", ".html"))
	app.StaticWeb("/static", "./web/static")

	session := sessions.New(sessions.Config{Cookie: "iris_session_id"})

	sparkApp := mvc.New(app)
	sparkApp.Register(session.Start)

	sparkApp.Handle(new(controllers.IndexController))
	sparkApp.Party("/spark").Configure(spark)

	app.Configure(iris.WithPostMaxMemory(maxSize)) //设置上传大小

	app.ConfigureHost(func(host *iris.Supervisor) { // <- 重要
		//您可以使用某些主机的方法控制流或延迟某些内容：
		// host.RegisterOnError
		// host.RegisterOnServe
		host.RegisterOnShutdown(func() {
			log.Info("Application shutdown on signal")
			task.GetTaskScheduleManager().Stop()
		})
	})

	return app
}

func main() {
	app := newApp()

	task.GetTaskScheduleManager().Start()

	app.Run(iris.Addr(":8080"))
}

func spark(app *mvc.Application) {
	sparkService := services.NewSparkService()
	app.Register(sparkService)

	app.Handle(new(controllers.SparkController))
}
