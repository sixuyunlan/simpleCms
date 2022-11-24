package user

import (
	"fmt"
	"simpleCms/app/common"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	type UserLogin struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	ul := &UserLogin{}
	if err := ctx.ShouldBindJSON(ul); err != nil {
		ctx.JSON(200, gin.H{
			"code":    2,
			"message": "参数错误",
		})
		return
	}
	user := &User{}
	if result := common.DB.Where("username = ?", ul.Username).First(user); result.Error != nil {
		ctx.JSON(200, gin.H{
			"code":    2,
			"message": "用户不存在",
		})
		return
	}
	if user.Password != password(ul.Password) {
		ctx.JSON(200, gin.H{
			"code":    2,
			"message": "密码错误",
		})
		return
	}
	claims := jwt.StandardClaims{
		Audience:  "SimplesSMS",
		ExpiresAt: time.Now().Add(60 * 24 * 3600 * time.Second).Unix(),
		Id:        fmt.Sprintf("%d", user.ID),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "SimplesSMS",
		NotBefore: 0,
		Subject:   "SimplesSMS-User-Auth",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(common.SecretKey))
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    2,
			"message": "token 生成错误",
		})
		return

	}

	ctx.JSON(200, gin.H{
		"code":    1,
		"message": "用户信息正确",
		"token":   tokenStr,
	})
	return
}
