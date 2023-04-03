package initializers

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

/* func awsURIMYSQL() string {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")

	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

	return dsn
} */

func ConnectToDB() /* Database */ {

	DB_USERNAME := "root"
	DB_PASSWORD := "genuine123"
	DB_HOST := "awsian.ct8srtp2hhm0.us-east-1.rds.amazonaws.com"
	DB_NAME := "aws"
	DB_PORT := "3306"

	var err error

	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

	/* dsn := awsURIMYSQL() */
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//no padrão ele faz db, err := cria o db e o erro ao mesmo tempo

	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados")
	}

	/* DB.AutoMigrate(&models.Post{}) */

}
