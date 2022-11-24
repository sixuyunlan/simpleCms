package user

import (
	"fmt"
	"net/http"
	"simpleCms/app/common"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Fuck(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    2,
			"message": "token不能为空",
		})
		ctx.Abort()

		return
	}
	token = token[7:]
	tokenObject, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(common.SecretKey), nil

	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    2,
			"message": "token错误",
		})
		ctx.Abort()

		return
	}

	if !tokenObject.Valid {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    2,
			"message": "token无效",
		})
		ctx.Abort()

		return

	}
	ctx.Next()

}
