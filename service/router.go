package service

import (
	"context"
	"log"
	"net/http"
	"nothing/middleware"
	"nothing/service/memory"
	"nothing/service/user"
	"os"
	"os/signal"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Run() {
	// gin with zap
	router := gin.New()
	logger, _ := zap.NewProduction()
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger, true))
	// 跨域
	router.Use(middleware.Cors())
	// 图片静态目录
	router.StaticFS("/images", http.Dir("./Data/images"))
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	// 用户
	users := router.Group("/user")
	{
		users.POST("/register", user.Register)
		users.POST("/login", user.Login)
		users.Use(middleware.AuthUser([]string{"user"}))
		users.GET("/info", user.GetUserInfo)
		users.PUT("/info", user.UpdateUserInfo)
		users.PUT("/password", user.ChangePassword)
		users.PUT("/flower/:uuid", user.Flower)
	}
	memorys := router.Group("/memory")
	memorys.Use(middleware.AuthUser([]string{"user"}))
	{
		memorys.GET("", memory.GetMyMemoryInfo)
		memorys.GET("/random", memory.GetRandomMemoryInfo)
		memorys.POST("", memory.AddMemory)
		memorys.DELETE("/:uuid", memory.DeleteMemory)
	}
	srv := &http.Server{
		Addr:    ":8001",
		Handler: router,
	}
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("关闭服务中 ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("服务已关闭:", err)
	}
	log.Println("服务已退出")
}
