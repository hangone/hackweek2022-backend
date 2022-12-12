package memory

import (
	"nothing/config"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMyMemoryInfo(c *gin.Context) {
	var memory []config.Memory
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "6"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "limit 错误",
		})
		return
	}
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "offset 错误",
		})
		return
	}
	username := c.GetString("username")
	data := config.Db.Where("creator = ?", username).Order("id desc").Select("uuid,title,content,name,created_at,updated_at").Find(&memory)
	total := data.RowsAffected
	result := data.Limit(limit).Offset(offset)
	if result.Error != nil {
		c.JSON(401, gin.H{
			"code":    401,
			"message": "获取失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    memory,
		"total":   total,
	})
}

func GetRandomMemoryInfo(c *gin.Context) {
	var memory []config.Memory
	var user config.User
	username := c.GetString("username")
	result := config.Db.Not("username = ?", username).Not("memory_count = 0").Order("random()").Take(&user)
	if result.Error != nil {
		c.JSON(401, gin.H{
			"code":    401,
			"message": "获取随机用户失败",
		})
		return
	}
	result = config.Db.Where("creator = ?", user.Username).Order("random()").Limit(6).Find(&memory)
	if result.Error != nil {
		c.JSON(401, gin.H{
			"code":    401,
			"message": "获取随机商品失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    memory,
	})
}
