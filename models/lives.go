package models

import (
	"errors"
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

func CreateLive(live *Live) error {
	lives, err := GetAllLives()
	if err != nil {
		return err
	}

	for _, l := range lives {
		if l.App == live.App {
			if l.Livename == live.Livename {
				return errors.New("live name is exist.")
			}
		}
	}

	err = orm.Gorm.Create(live).Error
	if err != nil {
		return err
	}

	return nil
}

func GetLiveById(id int) (*Live, error) {
	var lives []Live

	err := orm.Gorm.Where("id = ?", id).Find(&lives).Error
	if err != nil {
		return nil, err
	}

	if len(lives) == 0 {
		return nil, errors.New("error is rellay lives")
	}

	return &lives[0], err
}

func GetAllLives() ([]Live, error) {
	var lives []Live
	err := orm.Gorm.Find(&lives).Error
	if err != nil {
		return nil, err
	}
	return lives, nil
}

func DeleteLive(live *Live) error {
	err := orm.Gorm.Delete(live).Error
	if err != nil {
		return err
	}

	return nil
}

func CheckPublisherToken(appname string, livename string, token string) bool {
	var lives []Live
	err := orm.Gorm.Where("app = ?", appname).Where("livename = ?", livename).Where("publisher_token = ?", token).Find(&lives).Error

	if err != nil {
		log.Println(err)
		return false
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
		return false
	}

	if len(lives) == 1 {
		return true
	}

	return false
}

func CheckLive(livename string) bool {
	var lives []Live
	err := orm.Gorm.Where("livename = ?", livename).Find(&lives).Error

	if err != nil {
		log.Println(err)
		return false
	}

	if len(lives) == 1 {
		return true
	}

	return false
}
