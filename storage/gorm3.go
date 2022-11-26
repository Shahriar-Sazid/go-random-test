package storage

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Category struct {
	gorm.Model
	Title string `gorm:"type:varchar(255)"`
	Sort  int
}

type Contents struct {
	gorm.Model
	Category   Category
	CategoryID uint
	Title      string `gorm:"type:varchar(255)"`
	Content    string `gorm:"type:varchar(255)"`
}

func GormTest3() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("could not open database")
	}
	err = db.AutoMigrate(&Contents{}, &Category{})
	if err != nil {
		log.Fatal("could not migrate database")
	}
	createTestData3(db)
	fetchData3(db)
}

func createTestData3(db *gorm.DB) {
	category := Category{
		Title: "ABC",
		Sort:  1,
	}
	err := db.Create(&category).Error
	if err != nil {
		fmt.Println("failed to create user data")
	}

	content := Contents{
		CategoryID: category.ID,
		Title:      "Good Content Title",
		Content:    "Good Content",
	}

	err = db.Create(&content).Error
	if err != nil {
		fmt.Println("failed to create user data")
	}
}

func fetchData3(db *gorm.DB) {
	var cts []Contents
	if err := db.Find(&cts).Error; err != nil {
		fmt.Println("failed to load post")
	}
	fmt.Println(cts)
}
