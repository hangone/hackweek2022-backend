package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET, PUT, POST, DELETE, OPTIONS"},
		//AllowHeaders:     []string{"Origin, X-Requested-With, X-Extra-Header, Content-Type, Accept, Authorization"},
		//ExposeHeaders:    []string{"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"*"},
		//AllowCredentials: true,
		MaxAge: 480 * time.Hour,
	})
}
