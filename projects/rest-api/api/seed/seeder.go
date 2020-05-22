package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/victoralagwu/learn-go/projects/rest-api/api/models"
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
//Load :
func Load(db *gorm.DB)  {
	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("Cannot drop table: &v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("Cannot migrate table: %v", err)
	}
	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID
		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error

		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}