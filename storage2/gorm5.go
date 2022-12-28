package storage2

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Post struct {
	Body   string `gorm:"type:text"`
	Id     string `gorm:"type:uuid;primary_key"`
	Likes  []Like `gorm:"foreignkey:PostId;references:Id"`
	User   User   `gorm:"foreignkey:UserId;references:Id"`
	UserId string
}

type Like struct {
	Id     string `gorm:"type:uuid;primary_key"`
	PostId string `gorm:"type:uuid;not null"`
	User   User   `gorm:"foreignkey:UserId;references:Id"`
	UserId string
}

type User struct {
	Id   string `gorm:"type:uuid;primary_key"`
	Name string
}

func GormTest() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("could not open database")
	}
	err = db.AutoMigrate(&User{}, &Post{}, &Like{})
	if err != nil {
		log.Fatal("could not migrate database")
	}
	createTestData(db)
	fetchData(db)
}

func createTestData(db *gorm.DB) {
	users := []User{
		{Id: "0b83313d-1f85-4093-8621-efd2f21419d3", Name: "Shahriar"},
		{Id: "bddd6566-bcd2-4ad1-8eb9-65a23f5a9856", Name: "John"},
		{Id: "663c1328-dce2-4527-aecb-7fc478c229c2", Name: "Durand"}}
	err := db.Create(&users).Error
	if err != nil {
		log.Println("failed to create user data")
	}
	like := Like{
		Id:     "45ba45fc-0900-4fcc-80dd-c394170b777b",
		UserId: users[0].Id,
	}
	post := Post{
		Id:     "4cebb4c7-d44e-4160-a2df-a06f43211d45",
		Body:   "Test Post",
		Likes:  []Like{like},
		UserId: users[1].Id,
	}
	err = db.Create(&post).Error
	if err != nil {
		log.Println("failed to crete post")
	}
}

func fetchData(db *gorm.DB) {
	post := Post{
		Id: "4cebb4c7-d44e-4160-a2df-a06f43211d45",
	}
	if err := db.Preload("Likes").First(&post).Error; err != nil {
		log.Println("failed to load post")
	}
	log.Println(post)
}
