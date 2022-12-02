package database

import (
	"backend/models"
	mysql "backend/pkg/mysql"
	"fmt"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.Users{}, &models.EPresence{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed!")
	}

	fmt.Println("Migration Successful!")
}
