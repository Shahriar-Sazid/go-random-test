package storage

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Post struct {
	ID     string `gorm:"primarKey"`
	Title  string `gorm:"unique"`
	Body   string
	UserID string
	User   User `gorm:"foreignKey:UserID;references:ID"`
}

type User struct {
	ID       string `gorm:"primaryKey"`
	Email    string `gorm:"unique"`
	Name     string `gorm:"unique"`
	Password string
}

type PostInputForm struct {
	Title string
	Body  string
}

func GormTest4() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("could not open database")
	}

	err = db.AutoMigrate(&Post{}, &User{})
	if err != nil {
		log.Fatal("could not migrate database")
	}
	PostCreate(PostInputForm{Title: "ABC", Body: "Good Work!"}, User{ID: "98e86a24-1b01-4ef9-aaea-fd0181d128ea"}, db)
}

func PostCreate(data PostInputForm, user User, db *gorm.DB) (Post, error) {
	post := Post{Title: data.Title, Body: data.Body, UserID: user.ID}
	if err := db.Create(&post).Error; err != nil {
		return Post{}, err
	}

	log.Println("Post's User ID : ", post.UserID)
	log.Println("Post's User : ", post.User)

	return post, nil
}
