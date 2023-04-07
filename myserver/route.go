package main

import (
	"myserver/controller"
	"myserver/mainroot"
	"myserver/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	// info
	r.GET("/", mainroot.CheckExist)
	r.GET("/info", mainroot.CheckExist)
	r.GET("/ping", mainroot.PingTest)
	r.POST("/like_this", mainroot.LikeIt)
	// r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	// categoryRoutes := r.Group("/categories")
	// categoryController := controller.NewCategoryController()
	// categoryRoutes.POST("", categoryController.Create)
	// categoryRoutes.PUT("/:id", categoryController.Update) //替换
	// categoryRoutes.GET("/:id", categoryController.Show)
	// categoryRoutes.DELETE("/:id", categoryController.Delete)

	// postRoutes := r.Group("/posts")
	// postRoutes.Use(middleware.AuthMiddleware())
	// postController := controller.NewPostController()
	// postRoutes.POST("", postController.Create)
	// postRoutes.PUT("/:id", postController.Update) //替换
	// postRoutes.GET("/:id", postController.Show)
	// postRoutes.DELETE("/:id", postController.Delete)
	// postRoutes.POST("/page/list", postController.PageList)

	return r
}
