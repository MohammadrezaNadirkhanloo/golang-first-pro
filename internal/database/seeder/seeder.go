package seeder

import (
	"first-app/internal/modules/article/models"
	modelsUser "first-app/internal/modules/user/models"
	"first-app/pkg/database"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	db := database.Connections()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("secret"), 12)
	if err != nil {
		log.Fatal("hass password error")
	}
	user := modelsUser.User{Name: "ali", Email: "ss@g./ss", Password: string(hashedPassword)}

	db.Create(&user)

	log.Println("creat user success")

	for i := 1; i <= 10; i++ {
		article := models.Article{Title: fmt.Sprintf("random title %d", i), Content: fmt.Sprintf("random content %d", i), UserID: 1}
		db.Create(&article)
		log.Println("article success %d",i)
	}
	log.Println("Done success")
}
