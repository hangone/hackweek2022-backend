package product

import (
	"log"
	"nothing/config"

	"github.com/gin-gonic/gin"
)

func DeleteProduct(c *gin.Context) {
	var productBind, product config.Product
	err := c.ShouldBindJSON(&productBind)
	if err != nil {
		return
	}
	result := config.Db.Where("uid = ?", productBind.Uuid).Delete(&product)
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
