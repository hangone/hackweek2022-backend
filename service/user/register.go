package user

import (
	"log"
	"nothing/config"
	"nothing/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Register(c *gin.Context) {
	var userBind config.User
	if err := c.ShouldBindJSON(&userBind); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"code":    400,
			"message": "json 解析失败",
		})
		return
	}
	password := model.Encoding(userBind.Password)
	create := config.Db.Create(&config.User{
		Uuid:     uuid.New(),
		Username: userBind.Username,
		Password: password,
	})
	if create.Error != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "用户已存在",
		})
		return
	}
	c.JSON(201, gin.H{
		"code":    201,
		"message": "注册成功",
	})
}
