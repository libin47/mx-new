package main

import (
	"myserver/controller/category"
	"myserver/controller/mainroot"
	"myserver/controller/user"
	"myserver/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	// info
	r.GET("/", mainroot.CheckExist)
	r.GET("/info", mainroot.CheckExist)
	r.GET("/ping", mainroot.PingTest)
	r.GET("/like_this", mainroot.GetLikeCount)
	r.POST("/like_this", mainroot.LikeIt)
	r.GET("/clean_catch", mainroot.ClearCache)
	r.GET("/clean_redis", mainroot.ClearRedis)
	// category
	r.POST("/categories", category.NewCategory)
	// r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	r.POST("/api/auth/register", user.Register)
	r.POST("/api/auth/login", user.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), user.Info)

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
