package dao

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func init() {
	var err error
	userName := "root"
	password := "MySQL123.."
	host := "192.168.211.2"
	port := 3306
	dbName := "comic"
	charset := "utf8mb4"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True&charset=%s", userName, password, host, port, dbName, charset)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Println(err)
	} else {
		//成功
		// SqlSession.DB().SetMaxIdleConns(10)
		// SqlSession.DB().SetMaxOpenConns(100)
		// SqlSession.DB().SetConnMaxLifetime(time.Millisecond * 5)
		d, _ := Db.DB()
		d.SetMaxIdleConns(10)
		d.SetMaxOpenConns(100)
		d.SetConnMaxLifetime(time.Millisecond * 5)
		fmt.Println("数据库连接成功")
	}
}
