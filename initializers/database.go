package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {

	var err error

	dsn := os.Getenv("POSTGRES_URI")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	//no padrão ele faz db, err := cria o db e o erro ao mesmo tempo

	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados")
	}

	/* DB.AutoMigrate(&models.Post{}) */

}
