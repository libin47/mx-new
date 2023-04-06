package model

import (
	"github.com/qiniu/qmgo/field"
)

type User struct {
	field.DefaultField `bson:",inline"`
	Name               string `bson:"name"`
	Email              string `bson:"email"`
	Password           string `bson:"password"`
}
