package service

import (
	"log"
	"net/http"
	"nothing/middleware"
	"nothing/service/product"
	"nothing/service/user"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.StaticFS("/images", http.Dir("./Data/images"))
	users := r.Group("/user")
	{
		users.POST("/register", user.Register)
		users.POST("/login", user.Login)
		users.PUT("/password", middleware.AuthUser([]string{"user"}), user.ChangePassword)
	}
	products := r.Group("/product")
	products.Use(middleware.AuthUser([]string{"user"}))
	{
		products.GET("", product.GetProductInfo)
		products.POST("", product.AddProducts)
	}
	err := r.Run(":8001")
	if err != nil {
		log.Fatalf("listen: %s\n", err)
	}
}
