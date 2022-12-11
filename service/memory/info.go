package memory

import (
	"nothing/config"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMemoryInfo(c *gin.Context) {
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
	data := config.Db.Where("creator = ?", username).Order("created_at desc").Select("uuid,title,content,picture")
	result := data.Limit(limit).Offset(offset).Find(&memory)
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
		"total":   data.RowsAffected,
	})
}
