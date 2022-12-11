package memory

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"nothing/config"
	"path"

	"golang.org/x/crypto/sha3"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AddMemory(c *gin.Context) {
	filename := ""
	var allowType = map[string]string{"image/jpg": "", "image/jpeg": "", "image/png": "", "image/bmp": ""}
	file, header, err := c.Request.FormFile("file")
	title := c.PostForm("title")
	content := c.PostForm("content")
	username := c.GetString("username")
	uuidV4 := uuid.New()
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
		//sha256.Write(bit.Bytes())
		//hash := fmt.Sprintf("%x", sha256.Sum(nil))
		salt := []byte("nothing")
		sha256.Write(buf)
		sha256.Write(salt)
		hashWithSalt := fmt.Sprintf("%x", sha256.Sum(nil))[7:23]
		filename = hashWithSalt + path.Ext(header.Filename)
		create := config.Db.Create(&config.Picture{
			OriginName:   header.Filename,
			Name:         filename,
			HashWithSalt: hashWithSalt,
		})
		if create.Error != nil {
			log.Println(create.Error)
			c.JSON(201, gin.H{
				"code":    201,
				"message": "添加成功",
				"name":    filename,
				"uuid":    uuidV4,
			})
			return
		}
		if err := c.SaveUploadedFile(header, "./Data/images/"+filename); err != nil {
			log.Println(err)
			c.JSON(400, gin.H{
				"code":    400,
				"message": "图片保存失败",
			})
			return
		}
	}
	if result := config.Db.Create(&config.Memory{
		Creator: username,
		Uuid:    uuidV4,
		Title:   title,
		Content: content,
		Picture: filename,
	}); result.Error != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"code":    400,
			"message": "添加失败",
		})
		return
	}
	if filename != "" {
		c.JSON(200, gin.H{
			"code":     200,
			"message":  "添加成功",
			"filename": filename,
			"uuid":     uuidV4,
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "添加成功",
		"uuid":    uuidV4,
	})
}
