package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const projectDirName = "DEMO-PROJECT-BACKEND"

var instance *gorm.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1"
	dbname   = "mydb"
)

func LoadEnv() {
	projectNname := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectNname.Find([]byte(currentWorkDirectory))
	err := godotenv.Load(string(rootPath) + `.env`)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func GetConfigDatabase() *gorm.DB {

	LoadEnv()
	dbHost := os.Getenv("DB_HOST")
	Port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort, _ := strconv.ParseUint(Port, 10, 32)
	url := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUsername, dbPassword, dbName)
	db, err := gorm.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	instance = db
	return instance

}
