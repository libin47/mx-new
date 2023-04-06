package database

// 建立数据库连接
// 返回数据库操作模块
import (
	"context"
	"fmt"

	"github.com/qiniu/qmgo"
	"github.com/spf13/viper"
)

var DB *qmgo.Client
var UserCollection string = "user"

func InitDB(ctx context.Context) *qmgo.Client {
	cfg := qmgo.Config{
		Uri: viper.GetString("datasource.driverName") + "://" +
			viper.GetString("datasource.host") + ":" +
			viper.GetString("datasource.port"),
		Auth: &qmgo.Credential{
			AuthMechanism: "SCRAM-SHA-256",
			AuthSource:    viper.GetString("datasource.database"),
			Username:      viper.GetString("datasource.username"),
			Password:      viper.GetString("datasource.password"),
			PasswordSet:   true,
		},
	}
	client, err := qmgo.NewClient(ctx, &cfg)

	if err != nil {
		fmt.Println(err)
	}

	DB = client

	return client
}

func GetDB() *qmgo.Database {
	db := DB.Database(viper.GetString("datasource.database"))
	return db
}

func GetUserCollection() *qmgo.Collection {
	db := DB.Database(viper.GetString("datasource.database"))
	col := db.Collection(UserCollection)
	return col
}
