package models

import "go-live/orm"

type Live struct {
	Id             int
	App            string
	Livename       string
	PublisherToken string
	PlayerToken    string
}

func init() {
	orm.Gorm.AutoMigrate(new(Live))
}
