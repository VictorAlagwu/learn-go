package seed

import (
	"log"

	"github.com/victoralagwu/learn-go/projects/rest-api/models"
)

var users = []models.User{
	models.User{
		Name: "Tester",
		Email: "tester@gmail.com",
		Password: "password",
	},
	models.User{
		Name: "Demo",
		Email: "demo@gmail.com",
		Password: "password",
	},
}

var posts = []models.Post{
	models.Post{
		Title: "Title 1",
		Content: "Content of title one",
	},
	models.Post{
		Title: "Title 2",
		Content: "Content of title two",
	},
}

func Load(db *gorm.DB)  {
	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("Cannot drop table: &v")

	}
}