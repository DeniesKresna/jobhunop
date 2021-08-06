package Configs

import (
	"log"
	"os"
	"strconv"

	"github.com/DeniesKresna/jobhunop/Models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() (err error) {
	er := godotenv.Load(".env")
	if er != nil {
		log.Fatalf("Error loading .env file")
	}
	port, er := strconv.Atoi(os.Getenv("DB_PORT"))
	if er != nil {
		log.Fatal("Error Convert PORT")
		return
	}

	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + strconv.Itoa(port) + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	//log.Fatal(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB = db
	return err
}

func DatabaseMigrate() {
	DB.AutoMigrate(&Models.User{}, &Models.Role{}, &Models.Academy{})
}
