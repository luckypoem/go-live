package configure

import (
	"go-live/models"
	"log"
)

func CheckAppName(appname string) bool {
	lives, err := models.GetAppsByNameorLiveon(appname)

	if err != nil {
		log.Println(err)
	}

	if len(lives) == 1 {
		return true
	}

	return false
}

func GetStaticPushUrlList(appname string) ([]string, bool) {
	return nil, false
}
