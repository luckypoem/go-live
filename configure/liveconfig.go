package configure

import (
	"go-live/models"
)

func CheckAppName(appname string) bool {
	lives, _ := models.GetAllLives()

	for _, app := range lives {
		if (app.Appname == appname) && (app.Liveon == "on") {
			return true
		}
	}
	return false
}

func GetStaticPushUrlList(appname string) ([]string, bool) {
	return nil, false
}
