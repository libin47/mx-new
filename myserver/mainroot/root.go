package mainroot

import (
	"myserver/response"

	"github.com/gin-gonic/gin"
)

// check is exist
func CheckExist(ctx *gin.Context) {
	response.Success(ctx, gin.H{
		"name":     "bilin's web server",
		"author":   "Bilin <https://wind-watcher.cn>",
		"version":  "dev",
		"homepage": "https://github.com/libin47/mx-new",
	})
}

// ping
func PingTest(ctx *gin.Context) {
	response.Success(ctx, gin.H{
		"data": "pong",
	})
}

// like that
func LikeIt(ctx *gin.Context) {
	response.NoContent(ctx)
}

// get like num
func GetLikeCount(ctx *gin.Context) {
	response.NoContent(ctx)
}
