package user

import (
	"nothing/config"
	"nothing/model"

	"github.com/gin-gonic/gin"
)

type Password struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

func ChangePassword(c *gin.Context) {
	var passwordBind Password
	var user config.Users
	username, _ := c.Get("username")
	if err := c.ShouldBindJSON(&passwordBind); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "输入错误",
		})
		return
	}
	print(username)
	result := config.Db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": "用户不存在",
		})
		return
	}
	if model.Compare(user.Password, passwordBind.OldPassword) != nil {
		c.JSON(403, gin.H{
			"code":    403,
			"message": "旧密码错误",
		})
		return
	}
	newPassword := model.Encoding(passwordBind.NewPassword)
	result = config.Db.Where("username = ?", username).Updates(config.Users{Password: newPassword})
	if result.Error != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "修改失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "修改成功",
	})
}
