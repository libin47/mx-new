package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, httpStatus int, data gin.H) {
	ctx.JSON(httpStatus, data)
}

func Success(ctx *gin.Context, data gin.H) {
	Response(ctx, http.StatusOK, data)
}

func Fail(ctx *gin.Context, msg string, data gin.H) {
	Response(ctx, http.StatusOK, data)
}

func Unauthorized(ctx *gin.Context, data gin.H) {
	Response(ctx, http.StatusUnauthorized, data)
}

func NoContent(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{})
}
