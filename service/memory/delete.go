package memory

import (
	"log"
	"nothing/config"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func DeleteMemory(c *gin.Context) {
	uuid2, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"code":    400,
			"message": "uuid 输入错误",
		})
		return
	}
	result := config.Db.Delete(&config.Memory{}, uuid2)
	if result.Error != nil || result.RowsAffected != 1 {
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
