package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"time"
)

// DB is the underlying database connection
var DB *gorm.DB

func OpenConnection() {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	dialect := mysql.Open(connection)
	db, err := gorm.Open(dialect, &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Info),
		TranslateError: true,
	})

	if err != nil {
		fmt.Println("[DATABASE]::CONNECTION_ERROR")
		panic(err)
	}

	sqlDb, err := db.DB()

	if err != nil {
		panic(err)
	}

	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetConnMaxLifetime(30 * time.Minute)
	sqlDb.SetConnMaxLifetime(5 * time.Minute)

	DB = db
	fmt.Println("[DATABASE]::CONNECTED")
}

// migrate create -ext sql -dir database/migrations create_table_first
// migrate create -ext sql -dir database/migrations create_table_second
// migrate create -ext sql -dir database/migrations create_table_third

// migrate -database "mysql://root:password@tcp(localhost:3306)/url_shortener" -path database/migrations up
// migrate -database "mysql://root:password@tcp(localhost:3306)/url_shortener" -path database/migrations down

//migrate -database "mysql://root:password@tcp(localhost:3306)/url_shortener" -path database/migrations up 2
// migrate -database "mysql://root:password@tcp(localhost:3306)/url_shortener" -path database/migrations down 2
