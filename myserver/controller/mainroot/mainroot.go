package mainroot

import (
	"myserver/database"
	"myserver/middleware/response"
	"myserver/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// check is exist
func CheckExist(ctx *gin.Context) {
	// r := model.TestModel{Aaa: "a", Bbb: "b"}
	// r.CheckRepet(ctx, []string{"Aaa"})
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

// like this
func LikeIt(ctx *gin.Context) {
	println(ctx.ClientIP())
	num := database.GetRedisInt(ctx, "like_it", ctx.ClientIP())
	if num != 1 {
		result := database.SetRedisInt(ctx, "like_it", ctx.ClientIP(), 1)
		if !result {
			println("like count error in reids")
		}
		db := database.GetOptionsCollection()
		var likemodel model.LikeModel
		err := db.Find(ctx, bson.M{"name": "like"}).One(&likemodel)
		likemodel.Value += 1
		err = db.UpdateOne(ctx, bson.M{"name": "like"}, likemodel)
		if err != nil {
			println("like count err in mongo:", err)
		}
		response.NoContent(ctx)
	} else {
		response.Fail(ctx, gin.H{"ok": 0, "message": "一天一次就够啦"})
	}

}

// get like num
func GetLikeCount(ctx *gin.Context) {
	db := database.GetOptionsCollection()
	var likemodel model.LikeModel
	err := db.Find(ctx, bson.M{"name": "like"}).One(&likemodel)
	if err != nil {
		likemodel = model.LikeModel{
			Name:  "like",
			Value: 0,
		}
		_, err = db.InsertOne(ctx, likemodel)
		if err != nil {
			println("like count err in mongo:", err)
		}
	}
	response.Success(ctx, gin.H{"data": likemodel.Value})
}

// TODO: clear cache
func ClearCache(ctx *gin.Context) {
	response.NoContent(ctx)
}

// clear redis
func ClearRedis(ctx *gin.Context) {
	rdb := database.GetRDB(ctx)
	rdb.FlushDB(ctx)
	response.NoContent(ctx)
}
