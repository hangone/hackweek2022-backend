package product

import (
	"log"
	"nothing/config"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func AddProducts(c *gin.Context) {
	file, _ := c.FormFile("file")
	title := c.PostForm("title")
	content := c.PostForm("content")
	ext := strings.ToLower(path.Ext(file.Filename))
	if ext != ".jpg" && ext != ".png" {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "只支持上传jpg和png格式的图片",
		})
		return
	}
	filename := strconv.FormatInt(time.Now().UnixNano(), 10) + ".png"
	if err := c.SaveUploadedFile(file, "./Data/images/"+filename); err != nil {
		log.Print(err)
		c.JSON(400, gin.H{
			"code":    400,
			"message": "图片上传失败",
		})
		return
	}
	username := c.GetString("username")
	if result := config.Db.Create(&config.Products{
		Uuid:    uuid.New(),
		Title:   title,
		Content: content,
		Creator: username,
	}); result.Error != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "添加失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "添加成功",
		"data":    "/images/" + filename,
	})
}
