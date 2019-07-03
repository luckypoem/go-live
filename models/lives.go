package models

import (
	"go-live/orm"
	"log"
)

type Live struct {
	Id             int
	App            string `gorm:"not null"`
	Livename       string `gorm:"not null"`
	PublisherToken string `gorm:"not null;unique"`
	PlayerToken    string `gorm:"not null;unique"`
}

func init() {
	orm.Gorm.AutoMigrate(new(Live))
}

func CheckPublisherToken(appname string, livename string, token string) bool {
	var lives []Live
	err := orm.Gorm.Where("app = ?", appname).Where("livename = ?", livename).Where("publisher_token = ?", token).Find(&lives).Error

	if err != nil {
		log.Println(err)
	}

	if len(lives) == 1 {
		return true
	}

	return false
}

func CheckPlayerToken(appname string, livename string, token string) bool {
	var lives []Live
	err := orm.Gorm.Where("app = ?", appname).Where("livename = ?", livename).Where("player_token = ?", token).Find(&lives).Error

	if err != nil {
		log.Println(err)
	}

	if len(lives) == 1 {
		return true
	}

	return false
}
