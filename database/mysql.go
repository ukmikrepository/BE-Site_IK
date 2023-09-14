package database

import (
	"backend_ukmik/config"
	"fmt"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func ConnectionDB(config *config.ConfigEnv) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUsername, config.DBPassword, config.DBHost, config.DBName)

	// dsn := fmt.Sprintf("host=%s user=%s password=%s port=3306 dbname=%s sslmode=disable", config.DBHost, config.DBUsername, config.DBPassword, config.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("🚀 Failed to the Database, err message: ", err.Error())
		return nil
	}

	fmt.Println("🚀 Connected Successfully to the Database")
	return db
}
