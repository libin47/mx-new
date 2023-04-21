package category

import (
	"fmt"
	"myserver/controller"
	"myserver/database"
	"myserver/middleware/response"
	"myserver/model"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// category new
func NewCategory(ctx *gin.Context) {
	var requestCategory = model.CategoryModelRequest{}
	ctx.Bind(&requestCategory)
	var catype int
	if requestCategory.Type == "category" {
		catype = 0
	} else {
		catype = 1
	}
	categoryinsert := model.CategoryModelInsert{
		BaseModel: model.BaseModel{
			CreateAt: time.Now(),
			UpdateAt: time.Now(),
		},
		Name: requestCategory.Name,
		Slug: requestCategory.Slug,
		Type: catype,
	}

	// mongo
	db := database.GetCategoriesCollection()
	// dothis
	var data bson.M
	d, _ := bson.Marshal(categoryinsert)
	bson.Unmarshal(d, data)
	result := controller.InsertAfterCheck(ctx, data, bson.M{"name": requestCategory.Name}, db)
	fmt.Println(result)
}

// ping
func PingTest(ctx *gin.Context) {
	response.Success(ctx, gin.H{
		"data": "pong",
	})
}
