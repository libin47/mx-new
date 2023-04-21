package controller

import (
	"fmt"
	"myserver/middleware/response"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertAfterCheck(ctx *gin.Context, data bson.M, check_data bson.M, db *qmgo.Collection) bson.D {
	// check
	number, err := db.Find(ctx, bson.M{"$or": check_data}).Count()
	if err != nil {
		response.Fail(ctx, gin.H{"msg": "error when find mongodb"})
	}
	// if exist
	if number > 0 {
		response.Fail(ctx, gin.H{"msg": "the slug or category have existed!"})
	}
	// insert
	result, err := db.InsertOne(ctx, data)
	if err != nil {
		response.Fail(ctx, gin.H{"data": "errro when insert mongo"})
	}
	// find & return
	filter := bson.D{{"_id", result.InsertedID}}
	var resultreturn bson.D
	err = db.Find(ctx, filter).One(&resultreturn)
	fmt.Println(resultreturn)
	return resultreturn
}
