package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
}

type Post struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
	User    User   `json:"user"`
}

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "blog.db")
	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&User{}, &Post{})

	DB = database
}
