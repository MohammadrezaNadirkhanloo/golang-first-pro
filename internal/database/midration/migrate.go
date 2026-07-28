package midration

import (
	articleModels "first-app/internal/modules/article/models"
	userModels "first-app/internal/modules/user/models"
	"first-app/pkg/database"
	"fmt"
	"log"
)

func Migrate() {
	db := database.Connections()

	err :=db.AutoMigrate(&userModels.User{}, &articleModels.Article{})

	if err!=nil{
		log.Fatal("Cant migrate")
		return
	}

	fmt.Println("Migrate Done >>>")
}
