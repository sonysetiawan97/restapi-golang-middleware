package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// DB ...
var DB *gorm.DB

// DBConfig ...
type DBConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

// BuildDBConfig ...
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		Username: "root",
		Password: "password",
		DBName:   "golang-todos",
	}
	return &dbConfig
}

// DbURL ...
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
