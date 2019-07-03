package orm

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Gorm *gorm.DB

func init() {
	var err error

	Gorm, err = gorm.Open("mysql", "root:Xr111900@tcp(rm-wz9p9wn6719qplwioho.mysql.rds.aliyuncs.com:3306)/livego?charset=utf8")

	if err != nil {
		log.Fatal(err)
	}
}
