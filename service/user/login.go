package user

import (
	"nothing/config"
	"nothing/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var userBind, user config.User
	if err := c.ShouldBindJSON(&userBind); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "json 解析失败",
		})
		return
	}
	result := config.Db.Where("username = ?", userBind.Username).First(&user)
	if result.Error != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userBind.Password)) != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "用户名或密码错误",
		})
		return
	}
	signedToken, err := model.GenerateToken(user.Username)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "生成 Token 失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "登陆成功",
		"token":   signedToken,
	})
}
