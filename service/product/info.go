package product

import (
	"nothing/config"
	"strconv"
	"time"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

type products struct {
	Uuid    uuid.UUID
	Title   string
	Content string
	Picture string
	//Creator   string    `gorm:"not null" json:"creator"`
	Receiver  string
	IsSale    bool
	IsArchive bool
	CreatedAt time.Time
}

func GetProductInfo(c *gin.Context) {
	//var productList []config.Products
	var productList []products
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "6"))
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "输入错误",
		})
		return
	}
	username := c.GetString("username")
	result := config.Db.Order("id desc").Where("creator = ?", username).Limit(limit).Offset(offset).Find(&productList)
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
		"data":    productList,
		"total":   result.RowsAffected,
	})
}
