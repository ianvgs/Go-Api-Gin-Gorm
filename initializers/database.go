package initializers

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "ian"
const DB_PASSWORD = "1708"
const DB_NAME = "ian"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

/* // Database struct
type Database struct {
	DB *gorm.DB
} */

var DB *gorm.DB

func ConnectToDB() /* Database */ {

	/* 	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME") */

	var err error

	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//no padrão ele faz db, err := cria o db e o erro ao mesmo tempo

	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados")
	}

	/* DB.AutoMigrate(&models.Post{}) */

}
