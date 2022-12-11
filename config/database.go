package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

var Db = new(gorm.DB)

type Users struct {
	Username string `gorm:"not null;unique" json:"username" binding:"required"` // 用户名
	Password string `gorm:"not null" json:"password" binding:"required"`        // 密码
	Type     string `gorm:"not null;default:'user'" json:"type"`                // 用户类型，默认为 user，管理员为 admin
	ShopName string `json:"shopName"`                                           // 店铺名称
	BeLiked  int    `gorm:"not null;default:0" json:"BeLiked"`                  // 被点赞数
	gorm.Model
}

type Memory struct {
	Uuid      uuid.UUID `gorm:"not null;type:uuid;primarykey" json:"uuid"` // uuid
	Title     string    `json:"title"`                                     // 标题
	Content   string    `json:"content"`                                   // 内容
	Picture   string    `json:"picture"`                                   // 图片
	Creator   string    `gorm:"not null" json:"-"`                         // 创建者
	Receiver  string    `json:"-"`                                         // 接收者
	IsSale    bool      `gorm:"not null;default:false" json:"-"`           // 是否出售
	IsArchive bool      `gorm:"not null;default:false"  json:"-"`          // 是否归档
	//gorm.Model
	ID        uint      `gorm:"not null;sort:desc;autoIncrement" json:"-"` // id
	CreatedAt time.Time `json:"createdAt"`                                 // 创建时间
	UpdatedAt time.Time `json:"updatedAt"`                                 // 更新时间
	DeletedAt time.Time `gorm:"index" json:"-"`                            // 删除时间
}

type Picture struct {
	OriginName   string `json:"-"`               // 原始文件名
	Name         string `json:"name"`            // 文件名
	Hash         string `gorm:"unique" json:"-"` // 文件哈希
	HashWithSalt string `json:"-"`               // 带盐的哈希
}

func InitDb() {
	var (
		host     = Config.Database.Host
		port     = Config.Database.Port
		user     = Config.Database.User
		password = Config.Database.Password
		dbname   = Config.Database.Dbname
		sslmode  = Config.Database.SslMode
		TimeZone = Config.Database.TimeZone
		err      error
	)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s", host, port, user, password, dbname, sslmode, TimeZone)
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	err = Db.AutoMigrate(&Memory{}, &Users{}, &Picture{})
	if err != nil {
		log.Fatalln(err)
	}
}
