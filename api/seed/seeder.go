package seed

import (
	"fmt"
	"log"
	"time"

	"github.com/hcastellanos-dev/fullstack/api/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Name:           "Harim Castellanos",
		Email:          "hca@gmail.com",
		Password:       "password",
		LastName:       "Castellanos",
		SecondLastName: "Altamirano",
		BirthDate:      time.Now(), // Debes proporcionar un valor de tipo time.Time para BirthDate
		IsDriver:       false,
		PhoneNumber:    "123456789",
		Rol:            "Admin",
		IsActive:       true,
		Picture:        "",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	},
	models.User{
		Name:           "Mary Martinez",
		Email:          "mary@gmail.com",
		Password:       "password",
		LastName:       "Martinez",
		SecondLastName: "",
		BirthDate:      time.Now(), // Debes proporcionar un valor de tipo time.Time para BirthDate
		IsDriver:       true,
		PhoneNumber:    "987654321",
		Rol:            "Admin",
		IsActive:       true,
		Picture:        "",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "Title 1",
		Content: "Hello world 1",
	},
	models.Post{
		Title:   "Title 2",
		Content: "Hello world 2",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for _, user := range users {
		// Realiza acciones con el usuario, como validar, guardar en la base de datos, etc.
		// Por ejemplo, puedes validar el usuario antes de guardarlo
		err := user.Validate("create")
		if err != nil {
			log.Printf("Error al validar el usuario: %v", err)
			continue // Puedes optar por saltar el usuario en caso de error de validación
		}

		// Aquí puedes realizar otras acciones como guardar el usuario en la base de datos, etc.
		// Supongamos que tienes una instancia de la base de datos llamada "db"
		createdUser, err := user.SaveUser(db)
		if err != nil {
			log.Printf("Error al guardar el usuario: %v", err)
			continue // Puedes optar por saltar el usuario en caso de error al guardar
		}

		fmt.Printf("Usuario creado exitosamente: %v\n", createdUser)
	}
}
