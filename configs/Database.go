Package Configs

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
	var er error
	er = godotenv.Load()
	if er != nil {
		log.Fatal("Error getting env")
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("Error Convert PORT")
	}

	dsn := os.Getenv("DB_USER") + "@tcp(" + os.Getenv("DB_HOST") + ":" + strconv.Itoa(port) + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB = db
	return err
}

func DatabaseMigrate() {
	DB.AutoMigrate(&Models.User{})
}
