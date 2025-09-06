package seeders

import (
	"backend/configs"
	"backend/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func UserSeeder() {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin"), 12)
	if err != nil {
		log.Fatal("Failed to hashed password: ", err)
		return
	}

	users := []models.User{
		// admin
		{
			Email: "admin@admin.com",
			Password: string(hashedPassword),
			Name: "Administrator",
		},
	}

	for _, user := range users {
		result := configs.DB.FirstOrCreate(&user, models.User{Email: user.Email})
		if result.Error != nil {
			log.Printf("Seeder failed for user %s: %v", user.Email, result.Error)
		}
	}

	log.Println("User seeder is completed 🚀")
}