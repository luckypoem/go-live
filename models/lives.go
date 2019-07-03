package models

import "go-live/orm"

type Live struct {
	Id      int
	Appname string
	Liveon  string
	Hlson   string
	Token   string
}

func init() {
	orm.Gorm.AutoMigrate(new(Live))
}

func GetAllLives() ([]Live, error) {
	var lives []Live
	err := orm.Gorm.Find(&lives).Error
	if err != nil {
		return nil, err
	}
	return lives, nil
}
