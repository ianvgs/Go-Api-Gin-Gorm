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

	/* DB_USERNAME := "root"
	DB_PASSWORD := "genuine123"
	DB_HOST := "awsian.ct8srtp2hhm0.us-east-1.rds.amazonaws.com"
	DB_NAME := "aws"
	DB_PORT := "3306" */

	const DB_USERNAME = "roots"
	const DB_PASSWORD = ""
	const DB_NAME = "news_database"
	const DB_HOST = "127.0.0.1"
	const DB_PORT = "3306"

	var err error

	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

	/* dsn := awsURIMYSQL() */
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//Desse jeito ele não deixa a aplicação rodar
	/* if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados")
	} */

	if err != nil {
		log.Println("Erro ao conectar ao banco de dados:", err)
		// handle the error here, for example:
		// return an error message to the user
		// retry the connection after a certain amount of time
		/* By using log.Println instead of log.Fatal, the program will continue running even if the database connection fails. You can then handle the error in a way that makes sense for your application, such as showing an error message to the user or retrying the connection after a certain amount of time. */
	}

	/* DB.AutoMigrate(&models.Post{}) */

}
