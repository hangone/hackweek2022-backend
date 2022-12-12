package memory

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"nothing/config"
	"path"

	"github.com/gin-gonic/gin"
	uuid2 "github.com/google/uuid"
	"golang.org/x/crypto/sha3"
)

func AddMemory(c *gin.Context) {
	var allowType = map[string]string{"image/jpg": "", "image/jpeg": "", "image/png": "", "image/bmp": ""}
	file, header, err := c.Request.FormFile("file")
	title := c.PostForm("title")
	content := c.PostForm("content")
	username := c.GetString("username")
	uuidV4 := uuid2.New()
	hashWithSalt := ""
	hash := ""
	filename := ""
	if err == nil {
		buf := make([]byte, 512)
		_, err := file.Read(buf)
		if err != nil {
			log.Println(err)
			c.JSON(400, gin.H{
				"code":    400,
				"message": "文件类型解析失败",
			})
			return
		}
		contentType := http.DetectContentType(buf)
		if _, ok := allowType[contentType]; !ok {
			c.JSON(400, gin.H{
				"code":    400,
				"message": "只支持上传jpg、png和bmp格式的图片",
			})
			return
		}
		bit := bytes.NewBuffer(nil)
		if _, err = io.Copy(bit, file); err != nil {
			log.Println(err)
			c.JSON(400, gin.H{
				"code":    400,
				"message": "图片校验失败",
			})
			return
		}
		sha256 := sha3.New256()
		salt := []byte("nothing")
		sha256.Write(buf)
		hash = fmt.Sprintf("%x", sha256.Sum(nil))
		sha256.Write(salt)
		hashWithSalt = fmt.Sprintf("%x", sha256.Sum(nil))
		filename = hashWithSalt[7:23] + path.Ext(header.Filename)
		if err := c.SaveUploadedFile(header, "./Data/images/"+filename); err != nil {
			log.Println(err)
			c.JSON(400, gin.H{
				"code":    400,
				"message": "图片保存失败",
			})
			return
		}
		hash = fmt.Sprintf("%x", sha3.Sum256(bit.Bytes()))
	}
	if result := config.Db.Create(&config.Memory{
		Creator:      username,
		Uuid:         uuidV4,
		Title:        title,
		Content:      content,
		OriginName:   header.Filename,
		Name:         filename,
		Hash:         hash,
		HashWithSalt: hashWithSalt,
	}); result.Error != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"code":    400,
			"message": "添加失败",
		})
		return
	}
	var user config.User
	config.Db.Where("username = ?", username).First(&user)
	user.MemoryCount++
	config.Db.Save(&user)
	c.JSON(200, gin.H{
		"code":     200,
		"message":  "添加成功",
		"filename": filename,
		"uuid":     uuidV4,
	})
}
