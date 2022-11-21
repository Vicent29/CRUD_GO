package main

import (
	"CRUD-api/Config"
	"CRUD-api/Models"
	"CRUD-api/Routes"

	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})
	Config.DB.AutoMigrate(&Models.Table{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
