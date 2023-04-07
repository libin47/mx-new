package middleware

import (
	"fmt"
	"myserver/database"
	"myserver/jwttoken"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取 authorization header
		tokenString := ctx.GetHeader("Authorization")

		fmt.Print("请求token", tokenString)

		//validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:] //截取字符 Bearer [eyj.... 后面的部分

		token, claims, err := jwttoken.ParseToken(tokenString)

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		//token通过验证, 获取claims中的UserID
		email := claims.Email
		// fmt.Println("email:", email)
		name := claims.Name
		// fmt.Println("name:", name)

		DB := database.GetUserCollection()
		count, err := DB.Find(ctx, bson.M{"name": name, "email": email}).Count()
		if err != nil {
			fmt.Println(err.Error())
		}
		if count == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		//用户存在 将user信息写入上下文
		ctx.Set("name", name)
		ctx.Set("email", email)

		ctx.Next()
	}
}
