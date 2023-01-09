package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB

//Model is sample of common table structure
type Model struct {
	ID        uint       `gorm:"primary_key" json:"id,omitempty"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	//dbHost := os.Getenv("DB_HOST")
	//dbPort := os.Getenv("DB_PORT")

	msql := mysql.Config{}
	log.Println(msql)

	conn, err := gorm.Open("mysql", username+":"+password+"@/"+dbName+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("ran into error opening connection xo")
		fmt.Print(err)
	}
	db = conn

	//Printing query
	db.LogMode(true)

	//Automatically create migration as per model
	db.Debug().AutoMigrate(
		&User{},
	)
}

//GetDB function return the instance of db
func GetDB() *gorm.DB {
	return db
}
