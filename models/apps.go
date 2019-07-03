package models

import "go-live/orm"

type App struct {
	Id      int
	Appname string `gorm:"not null;unique"`
	Liveon  string `gorm:"not null"`
}

func init() {
	orm.Gorm.AutoMigrate(new(App))
}

func CreateApp(app *App) error {
	err := orm.Gorm.Create(app).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAllApps() ([]App, error) {
	var apps []App
	err := orm.Gorm.Find(&apps).Error
	if err != nil {
		return nil, err
	}
	return apps, nil
}

func GetAppsByNameorLiveon(appname string) ([]App, error) {
	var apps []App

	err := orm.Gorm.Where("appname = ?", appname).Where("liveon = ?", "on").Find(&apps).Error

	if err != nil {
		return nil, err
	}

	return apps, nil
}
