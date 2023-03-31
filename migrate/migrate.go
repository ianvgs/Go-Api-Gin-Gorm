package main

import (
	"goagain/initializers"
	"goagain/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	initializers.DB.AutoMigrate(&models.Colaborador{}, &models.Noticia{}, &models.Categoria{})

}
