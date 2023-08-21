package orm

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// GetDB 데이터베이스 연결
func GetDB(dbName string) *gorm.DB {
	var err error

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: 200 * time.Millisecond, // Slow SQL threshold. Set to 0 to disable.
			LogLevel:      logger.Silent,          // Log level
			Colorful:      false,                  // Disable color
		},
	)

	db := &gorm.DB{}
	if dbName == "postgres" {
		dsn := "host=127.0.0.1 user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Seoul"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	} else if dbName == "mysql" {
		dsn := "root:mysql@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	} else if dbName == "sqlite" {
		db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{Logger: newLogger})
	} else {
		panic("invalid db name")
	}
	if err != nil {
		panic(err)
	}
	return db
}
