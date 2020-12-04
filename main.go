package main

import (
	"fmt"

	"example.com/v1/config"
	"example.com/v1/models"
	"example.com/v1/route"
	"github.com/jinzhu/gorm"
)

// err ...
var err error

// main ...
func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("status: ", err)
	}

	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Todo{})

	r := route.Routes()
	r.Run()
}
