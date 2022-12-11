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
	Username string `gorm:"not null;unique" json:"username" binding:"required"`
	Password string `gorm:"not null" json:"password" binding:"required"`
	Type     string `gorm:"not null;default:'user'" json:"type"` // 默认用户类型为 user
	ShopName string `json:"shopName"`
	BeLiked  int    `gorm:"not null;default:0" json:"BeLiked"`
	gorm.Model
}

type Product struct {
	Uuid      uuid.UUID `gorm:"not null;type:uuid;primarykey" json:"uuid"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Picture   string    `json:"picture"`
	Creator   string    `gorm:"not null" json:"-"`
	Receiver  string    `json:"-"`
	IsSale    bool      `gorm:"not null;default:false" json:"-"`
	IsArchive bool      `gorm:"not null;default:false"  json:"-"`
	//gorm.Model
	ID        uint      `gorm:"not null;sort:desc;autoIncrement" json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `gorm:"index" json:"-"`
}

type Picture struct {
	OriginName   string `json:"-"`
	Name         string `json:"name"`
	Hash         string `json:"-"`
	HashWithSalt string `json:"-"`
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
	err = Db.AutoMigrate(&Product{}, &Users{}, &Picture{})
	if err != nil {
		log.Fatalln(err)
	}
}
