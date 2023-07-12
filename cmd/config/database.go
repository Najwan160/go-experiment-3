package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func ConnectDB() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		Env.DB.Username,
		Env.DB.Password,
		Env.DB.Host,
		Env.DB.Port,
		Env.DB.Schema,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		CreateBatchSize:        Env.DB.CreateBatchSize,
	})
	if err != nil {
		panic(fmt.Sprintf("unable to open connection to database: %v", err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("unable get database instance: %v", err))
	}

	if err := sqlDB.Ping(); err != nil {
		panic(fmt.Sprintf("failed to ping database: %v", err))
	}

	if Env.DB.Debug {
		db = db.Debug()
	}

	DB = db
}
