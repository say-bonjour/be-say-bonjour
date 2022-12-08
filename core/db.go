package core

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabase(properties DatabaseProperties) *gorm.DB {
	dsn := properties.Username + ":" + properties.Password + "@" + properties.Url + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err.Error())
	}

	if properties.PoolSize > 0 {
		sqlDB.SetMaxIdleConns(properties.PoolSize)
		sqlDB.SetMaxOpenConns(properties.PoolSize)
	} else {
		sqlDB.SetMaxIdleConns(2)
		sqlDB.SetMaxOpenConns(2)
	}

	return db
}
