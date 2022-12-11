package memory

import (
	"log"
	"nothing/config"

	"github.com/gin-gonic/gin"
)

func DeleteMemory(c *gin.Context) {
	var memoryBind, memory config.Memory
	err := c.ShouldBindJSON(&memoryBind)
	if err != nil {
		return
	}
	result := config.Db.Where("uid = ?", memoryBind.Uuid).Delete(&memory)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(400, gin.H{
			"code":    400,
			"message": "删除失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "删除成功",
	})
}
