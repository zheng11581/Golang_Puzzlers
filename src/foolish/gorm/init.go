package main

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var DSN = "L/vnpwfB7po=:+5KOFFuKHJjA+nXPXNgXspcD22DXBT2C@tcp(rm-2ze1150923152oc28.mysql.rds.aliyuncs.com:3306)/gw_manage?useUnicode=true&amp;characterEncoding=UTF-8"

func init() {
	var err error
	jdbcUrl, err := url.QueryUnescape(DSN)
	if err != nil {
		return
	}

	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               jdbcUrl, // data source name
		DefaultStringSize: 256,     // default size for string fields

	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Printf("GORM init error: %v", err)
		return
	}
	setPool(DB)
}

func setPool(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

}

func main() {
	fmt.Println(DB.Name())
}
