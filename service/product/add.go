package product

import (
	"fmt"
	"log"
	"math/rand"
	"nothing/config"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AddProducts(c *gin.Context) {
	file, err := c.FormFile("file")
	filename := ""
	if err == nil {
		ext := strings.ToLower(path.Ext(file.Filename))
		if ext != ".jpg" && ext != ".png" {
			c.JSON(400, gin.H{
				"code":    400,
				"message": "只支持上传jpg和png格式的图片",
			})
			return
		}
		rand.Seed(time.Now().UnixNano())
		randNumber := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
		filename = strconv.FormatInt(time.Now().UnixNano(), 10) + randNumber + path.Ext(file.Filename)
		if err := c.SaveUploadedFile(file, "./Data/images/"+filename); err != nil {
			log.Print(err)
			c.JSON(400, gin.H{
				"code":    400,
				"message": "图片上传失败",
			})
			return
		}
	}
	title := c.PostForm("title")
	content := c.PostForm("content")
	username := c.GetString("username")
	if result := config.Db.Create(&config.Products{
		Uuid:    uuid.New(),
		Title:   title,
		Content: content,
		Picture: filename,
		Creator: username,
	}); result.Error != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "添加商品失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "添加成功",
		"data":    "/images/" + filename,
	})
}
