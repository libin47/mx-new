package controller

import (
	"fmt"
	"myserver/database"
	"myserver/jwttoken"
	"myserver/model"
	"myserver/response"

	"net/http"
	"regexp"

	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func IsEmail(email string) bool {
	reg := regexp.MustCompile(`\w+@\w+\.com`)
	result := reg.FindAllString(email, -1)
	return len(result) != 0
}

// 检查数据库中是否存在name
func IsExistName(ctx *gin.Context, name string, db *qmgo.Collection) bool {
	count, err := db.Find(ctx, bson.M{"name": name}).Count()
	if err != nil {
		fmt.Println(err.Error())
	}
	return count > 0
}

// 检查数据库中是否存在email
func IsExistEmail(ctx *gin.Context, email string, db *qmgo.Collection) bool {
	count, err := db.Find(ctx, bson.M{"email": email}).Count()
	if err != nil {
		fmt.Println(err.Error())
	}
	return count > 0
}

func Register(ctx *gin.Context) {
	DB := database.GetUserCollection()
	//使用map获取请求参数
	var requestUser = model.User{}
	ctx.Bind(&requestUser)
	//获取参数
	fmt.Print()
	name := requestUser.Name
	email := requestUser.Email
	password := requestUser.Password
	// 数据验证
	// 用户名长度
	if len(name) > 18 || len(name) < 3 {
		response.Response(ctx, http.StatusUnprocessableEntity, gin.H{"msg": "用户名长度应在2-18之间"})
		return
	}
	// 密码长度
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, gin.H{"msg": "密码不能少于6位"})
		return
	}
	// 邮箱格式
	if !IsEmail(email) {
		response.Response(ctx, http.StatusUnprocessableEntity, gin.H{"msg": "邮箱格式错误，请输入正确邮箱"})
		return
	}
	// 检测用户名是否存在
	if IsExistName(ctx, name, DB) {
		response.Response(ctx, http.StatusUnprocessableEntity, gin.H{"msg": "该用户名已存在，请更换用户名"})
		return
	}
	// 邮箱是否存在
	if IsExistEmail(ctx, email, DB) {
		response.Response(ctx, http.StatusUnprocessableEntity, gin.H{"msg": "该邮箱已注册，请直接登录"})
		return
	}
	//创建用户
	hasePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, gin.H{"msg": "加密失败"})
		return
	}
	newUser := &model.User{
		Name:     name,
		Email:    email,
		Password: string(hasePassword),
	}
	_, err = DB.InsertOne(ctx, newUser)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, gin.H{"msg": err.Error()})
		return
	}

	//返回结果
	//发放token
	token, err := jwttoken.ReleaseToken(newUser)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, gin.H{"msg": "系统异常"})
		fmt.Printf("token generate error: %v", err)
		return
	}

	//返回结果
	response.Success(ctx, gin.H{"token": token, "msg": "注册成功"})
}

func Login(ctx *gin.Context) {
	DB := database.GetUserCollection()
	//获取数据
	//使用map获取请求参数
	// fmt.Println(ctx)
	var requestUser = model.User{}
	ctx.Bind(&requestUser)
	//获取参数
	name := requestUser.Name
	password := requestUser.Password
	//数据验证-name是否是用户名或邮箱
	var user model.User
	if !IsExistName(ctx, name, DB) && !IsExistEmail(ctx, name, DB) {
		response.Response(ctx, http.StatusUnprocessableEntity, gin.H{"msg": "用户不存在"})
		return
	}
	if IsExistName(ctx, name, DB) {
		DB.Find(ctx, bson.M{"name": name}).One(&user)
	} else {
		DB.Find(ctx, bson.M{"email": name}).One(&user)
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusBadRequest, gin.H{"msg": "密码错误"})
		return
	}

	//发放token
	token, err := jwttoken.ReleaseToken(&user)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, gin.H{"msg": "系统异常"})
		return
	}

	//返回结果
	response.Success(ctx, gin.H{"token": token})
}

type UserDto struct {
	Name  string `bson:"name"`
	Email string `bson:"email"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:  user.Name,
		Email: user.Email,
	}
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("name")
	fmt.Print(user)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"user": user},
	})
}
