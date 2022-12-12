package middleware

import (
	"nothing/config"
	"nothing/model"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthUser(userTypes []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		parts := strings.SplitN(token, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(498, gin.H{
				"code":    498,
				"message": "Authorization 格式错误",
			})
			c.Abort()
			return
		}
		claims, err := model.ParseToken(parts[1])
		if err != nil {
			c.JSON(498, gin.H{
				"code":    498,
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		var user config.User
		result := config.Db.Where("username = ?", claims.Username).First(&user)
		if result.Error != nil {
			c.JSON(401, gin.H{
				"code":    401,
				"message": "用户不存在",
			})
			c.Abort()
			return
		}
		for _, userType := range userTypes {
			if user.Type == userType {
				c.Set("uuid", user.Uuid)
				c.Set("username", user.Username)
				c.Next()
				c.Abort()
				return
			}
		}
	}
}
