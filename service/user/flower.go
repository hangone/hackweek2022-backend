package user

import (
	"log"
	"nothing/config"

	"github.com/gin-gonic/gin"
	uuid2 "github.com/google/uuid"
)

func Flower(c *gin.Context) {
	uuid, err := uuid2.Parse(c.Param("uuid"))
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"code":    400,
			"message": "uuid 输入错误",
		})
		return
	}
	var user config.User
	result := config.Db.First(&user, uuid).Update("flower", user.Flower+1)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(400, gin.H{
			"code":    400,
			"message": "送花失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "送花成功",
	})
}
