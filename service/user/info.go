package user

import (
	"log"
	"nothing/config"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	var user config.User
	username := c.GetString("username")
	result := config.Db.Where("username = ?", username).Select("uuid,username,type,shop_name,liked,flower,created_at,updated_at").First(&user)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(401, gin.H{
			"code":    401,
			"message": "获取用户信息失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    user,
	})
}

func UpdateUserInfo(c *gin.Context) {
	var user, userBind config.User
	if err := c.ShouldBind(&userBind); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"code":    400,
			"message": "json 解析失败",
		})
		return
	}
	username := c.GetString("username")
	result := config.Db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(401, gin.H{
			"code":    401,
			"message": "获取用户信息失败",
		})
		return
	}
	result.Updates(&config.User{
		ShopName: userBind.ShopName,
		Liked:    userBind.Liked,
	})
	c.JSON(200, gin.H{
		"code":    200,
		"message": "更新成功",
	})
}
