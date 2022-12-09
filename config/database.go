package config

import (
	"fmt"
	"log"

	"github.com/google/uuid"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db = new(gorm.DB)

type Users struct {
	Username string `gorm:"not null;unique" json:"username" binding:"required"`
	Password string `gorm:"not null" json:"password" binding:"required"`
	Type     string `gorm:"not null;default:'user'" json:"type"`
	gorm.Model
}

type Products struct {
	Uuid      uuid.UUID `gorm:"not null;unique;type:uuid" json:"uuid"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Picture   string    `json:"picture"`
	Creator   string    `gorm:"not null" json:"creator"`
	Receiver  string    `json:"receiver"`
	IsSale    bool      `gorm:"not null;default:false" json:"isSale"`
	IsArchive bool      `gorm:"not null;default:false" json:"isArchive"`
	gorm.Model
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
	err = Db.AutoMigrate(&Products{}, &Users{})
	if err != nil {
		log.Fatalln(err)
	}
}
