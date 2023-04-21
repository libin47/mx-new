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

// collection name
var AlbumsCollection string = "albums"
var AnalyzesCollection string = "analyzes"
var CategoriesCollection string = "categories"
var CommentsCollection string = "comments"
var LinksCollection string = "links"
var NotesCollection string = "notes"
var OptionsCollection string = "options"
var PagesCollection string = "pages"
var PhotosCollection string = "photos"
var PostsCollection string = "posts"
var ProjectsCollection string = "projects"
var QasCollection string = "qas"
var RecentliesCollection string = "recentlies"
var SaysCollection string = "says"
var ServerlessstoragesCollection string = "serverlessstorages"
var SnippetsCollection string = "snippets"
var TopicsCollection string = "topics"
var UsersCollection string = "users"

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

func GetAlbumsCollection() *qmgo.Collection {
	db := DB.Database(viper.GetString("datasource.database"))
	col := db.Collection(AlbumsCollection)
	return col
}

func GetAnalyzesCollection() *qmgo.Collection {
	db := DB.Database(viper.GetString("datasource.database"))
	col := db.Collection(AnalyzesCollection)
	return col
}
func GetCategoriesCollection() *qmgo.Collection {
	db := DB.Database(viper.GetString("datasource.database"))
	col := db.Collection(CategoriesCollection)
	return col
}

func GetCommentsCollection() *qmgo.Collection {
	db := DB.Database(viper.GetString("datasource.database"))
	col := db.Collection(CommentsCollection)
	return col
}

func GetLinksCollection() *qmgo.Collection {
	db := DB.Database(viper.GetString("datasource.database"))
	col := db.Collection(LinksCollection)
	return col
}

func GetNotesCollection() *qmgo.Collection {
	db := DB.Database(viper.GetString("datasource.database"))
	col := db.Collection(NotesCollection)
	return col
}

func GetOptionsCollection() *qmgo.Collection {
	db := DB.Database(viper.GetString("datasource.database"))
	col := db.Collection(OptionsCollection)
	return col
}

func GetPagesCollection() *qmgo.Collection {
	db := DB.Database(viper.GetString("datasource.database"))
	col := db.Collection(PagesCollection)
	return col
}

func GetPostsCollection() *qmgo.Collection {
	db := DB.Database(viper.GetString("datasource.database"))
	col := db.Collection(PostsCollection)
	return col
}

func GetPhotosCollection() *qmgo.Collection {
	db := DB.Database(viper.GetString("datasource.database"))
	col := db.Collection(PhotosCollection)
	return col
}

func GetProjectsCollection() *qmgo.Collection {
	db := DB.Database(viper.GetString("datasource.database"))
	col := db.Collection(ProjectsCollection)
	return col
}

func GetQasCollection() *qmgo.Collection {
	db := DB.Database(viper.GetString("datasource.database"))
	col := db.Collection(QasCollection)
	return col
}

func GetRecentliesCollection() *qmgo.Collection {
	db := DB.Database(viper.GetString("datasource.database"))
	col := db.Collection(RecentliesCollection)
	return col
}

func GetSaysCollection() *qmgo.Collection {
	db := DB.Database(viper.GetString("datasource.database"))
	col := db.Collection(SaysCollection)
	return col
}

func GetServerlessstoragesCollection() *qmgo.Collection {
	db := DB.Database(viper.GetString("datasource.database"))
	col := db.Collection(ServerlessstoragesCollection)
	return col
}

func GetSnippetsCollection() *qmgo.Collection {
	db := DB.Database(viper.GetString("datasource.database"))
	col := db.Collection(SnippetsCollection)
	return col
}

func GetTopicsCollection() *qmgo.Collection {
	db := DB.Database(viper.GetString("datasource.database"))
	col := db.Collection(TopicsCollection)
	return col
}

func GetUserCollection() *qmgo.Collection {
	db := DB.Database(viper.GetString("datasource.database"))
	col := db.Collection(UsersCollection)
	return col
}
