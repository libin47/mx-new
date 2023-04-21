package model

import (
	"fmt"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/qmgo"
)

type BaseModel struct {
	CreateAt time.Time `bson:"createAt"`
	UpdateAt time.Time `bson:"updateAt"`
}
type TestModel struct {
	Aaa string `bson:"aaa"`
	Bbb string `bson:"bbb"`
}

type BaseStore interface {
}

func CheckRepetBase(ctx *gin.Context, model BaseStore, id []string, db *qmgo.Collection) bool {
	var quary map[string]string
	value := reflect.ValueOf(model)
	for _, id_ := range id {
		v := value.FieldByName(id_)
		fmt.Println("value:", value.FieldByName(id_))
		quary[id_] = v.String()
	}
	// db.Find(bson.D{quary})
	return true
}
