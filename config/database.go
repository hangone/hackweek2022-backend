package config

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db = new(gorm.DB)

type GormModel struct {
	//gorm.Model
	ID        uint           `gorm:"not null;autoIncrement" json:"-"` // id
	CreatedAt time.Time      `json:"createdAt"`                       // 创建时间
	UpdatedAt time.Time      `json:"updatedAt"`                       // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                  // 删除时间
}
type User struct {
	Uuid        uuid.UUID `gorm:"not null;type:uuid;primaryKey" json:"uuid"`
	Username    string    `gorm:"not null;unique" json:"username"`       // 用户名
	Password    string    `gorm:"not null,-" json:"password,omitempty"`  // 密码
	Type        string    `gorm:"not null;default:'user'" json:"type"`   // 用户类型，默认为 user，管理员为 admin
	ShopName    string    `json:"shopName"`                              // 店铺名称
	MemoryCount int       `gorm:"not null;default:0" json:"memoryCount"` // 记忆总数
	Liked       string    `json:"liked"`                                 // 兴趣
	Flower      int       `gorm:"not null;default:0" json:"flower"`      // 花
	GormModel
}

type Memory struct {
	Uuid         uuid.UUID `gorm:"not null;type:uuid;primarykey" json:"uuid"` // uuid
	Title        string    `json:"title"`                                     // 标题
	Content      string    `json:"content"`                                   // 内容
	OriginName   string    `json:"-"`                                         // 图片原始文件名
	Name         string    `json:"name"`                                      // 图片文件名
	Hash         string    `json:"-"`                                         // 图片文件哈希
	HashWithSalt string    `json:"-"`                                         // 图片带盐的哈希
	Creator      string    `gorm:"not null" json:"-"`                         // 创建者
	Receiver     string    `json:"-"`                                         // 接收者
	IsSale       bool      `gorm:"not null;default:false" json:"-"`           // 是否出售
	IsArchive    bool      `gorm:"not null;default:false" json:"-"`           // 是否归档
	GormModel
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
	err = Db.AutoMigrate(&Memory{}, &User{})
	if err != nil {
		log.Fatalln(err)
	}
}
