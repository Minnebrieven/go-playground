package configs

import (
	"fmt"

	"swim-class/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	DB_Host     string
	DB_Port     int
	DB_Name     string
	DB_Username string
	DB_Password string
}

func ConnectDB() (*gorm.DB, error) {
	config := Config{
		DB_Host:     "localhost",
		DB_Port:     3306,
		DB_Name:     "swim_class",
		DB_Username: "root",
		DB_Password: "",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)
	return gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		models.User{},
		models.Class{},
		models.Instructor{},
		models.Appointment{},
	)
}
