package initializer

import (
	"atm/api"
	"atm/middleware"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	engine := gin.Default()
	// 中间件的注册时机也有讲究，某中间件注册之前的API不会走这个中间件
	engine.Use(middleware.Cors())

	// I think it is a bad idea to design these RESTful APIs:
	// users := engine.Group("/api/v1/user")
	// {
	// 	users.GET("/info", )
	// 	users.POST("/info", )
	// 	users.PUT("/info", )
	// 	users.DELETE("/info", )

	// 	users.GET("/verifycode", )

	// 	users.POST("/login", )
	// 	users.POST("/newpass", )
	// }

	// students := engine.Group("/api/v1/student")
	// {
	// 	students.GET("/list", )
	// 	students.GET("/option", )

	// 	students.GET("/info", )
	// 	students.POST("/info", )
	// 	students.PUT("/info", )
	// 	students.DELETE("/info", )

	// 	students.POST("/email", )
	// 	students.GET("/excel", )
	// }

	// projects := engine.Group("/api/v1/project")
	// {
	// 	projects.GET("/list")

	// 	projects.GET("/info", )
	// 	projects.POST("/info", )
	// 	projects.PUT("/info", )
	// 	projects.DELETE("/info", )

	// 	projects.GET("/excel", )
	// }

	// outputs := engine.Group("/api/v1/output")
	// {
	// 	outputs.GET("/list")

	// 	outputs.GET("/info", )
	// 	outputs.POST("/info", )
	// 	outputs.PUT("/info", )
	// 	outputs.DELETE("/info", )

	// 	outputs.GET("/excel", )
	// }

	// utils := engine.Group("/api/v1/util")
	// {
	// 	utils.GET("/dbstate", )
	// 	utils.POST("/file", )
	// 	utils.DELETE("/file", )

	// }

	// summary := engine.Group("/api/v1/summary")
	// {
	// 	summary.GET("/", )
	// }

	// mailconf := engine.Group("/api/v1/mailconf")
	// {
	// 	mailconf.GET("/status", )
	// 	mailconf.PUT("/status", )

	// 	mailconf.GET("/info", )
	// 	mailconf.PUT("/info", )
	// 	mailconf.DELETE("/info", )
	// }

	users := engine.Group("/api/v1/user")
	{
		users.GET("/verifycode", api.NewUserAPI().GetVerifyCode)
		users.POST("/login", api.NewUserAPI().Login)
		users.POST("/register", api.NewUserAPI().Register)
		users.POST("/forgetpass", api.NewUserAPI().ForgetPass)

		engine.Use(middleware.JwtAuth())

		// 以下操作以及后面对操作都要登陆了才能做
		users.GET("/info", api.NewUserAPI().GetInfo)
		// get原则上没有禁止带body，但是gin不会管get里面的body
		users.POST("/delete", api.NewUserAPI().Delete)
	}

	students := engine.Group("/api/v1/student")
	{
		students.GET("/list", api.NewStudentAPI().GetList)
		students.GET("/option", api.NewStudentAPI().GetOption)
		students.GET("/info", api.NewStudentAPI().GetInfo)
		students.GET("/export", api.NewStudentAPI().ExportList)

		students.POST("/create", api.NewStudentAPI().CreateInfo)
		students.POST("/update", api.NewStudentAPI().UpdateInfo)
		students.POST("/delete", api.NewStudentAPI().Delete)

		students.POST("/send", api.NewStudentAPI().SendEmail)
	}

	projects := engine.Group("/api/v1/project")
	{
		projects.GET("/list", api.NewProjectAPI().GetList)
		projects.GET("/info", api.NewProjectAPI().GetInfo)
		projects.GET("/export", api.NewProjectAPI().ExportList)
		
		projects.POST("/outputs/update", api.NewProjectAPI().UpdateOutputs)
		projects.POST("/create", api.NewProjectAPI().CreateInfo)
		projects.POST("/update", api.NewProjectAPI().UpdateInfo)
		projects.POST("/delete", api.NewProjectAPI().Delete)
	}

	outputs := engine.Group("/api/v1/output")
	{
		outputs.GET("/list", api.NewOutputAPI().GetList)
		outputs.GET("/info", api.NewOutputAPI().GetInfo)
		outputs.GET("/export", api.NewOutputAPI().ExportList)

		outputs.POST("/create", api.NewOutputAPI().CreateInfo)
		outputs.POST("/update", api.NewOutputAPI().UpdateInfo)
		outputs.POST("/delete", api.NewOutputAPI().Delete)
	}

	utils := engine.Group("/api/v1/util")
	{
		utils.GET("/initdb", api.NewUtilAPI().InitDB)

		utils.POST("/uploadfile", api.NewUtilAPI().UploadFile)
		utils.POST("/removefile", api.NewUtilAPI().RemoveFile)
	}

	summary := engine.Group("/api/v1/summary")
	{
		summary.GET("", api.NewSummaryAPI().Get)
	}

	mailconf := engine.Group("/api/v1/mailconf")
	{
		mailconf.POST("/switch/set", api.NewMailConfAPI().SetSwitch)
		mailconf.GET("/info", api.NewMailConfAPI().GetInfo)
		mailconf.GET("/check", api.NewMailConfAPI().CheckValidity)

		mailconf.POST("/save", api.NewMailConfAPI().Save)
		mailconf.POST("/delete", api.NewMailConfAPI().Delete)
	}

	return engine
}
