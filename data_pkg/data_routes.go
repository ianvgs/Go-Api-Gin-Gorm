package data_routes

import (
	examplepkg "goagain/data_pkg/controllers"
	data_controller "goagain/data_pkg/controllers/data_controller"
	sentiment_controller "goagain/data_pkg/controllers/sentiment_controller"

	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func DataRoutes(g *gin.RouterGroup) {

	//data_pkg/controllers/sayhi.go
	g.GET("/sayhi", examplepkg.SayHi())
	//data_pkg/controllers/saybye.go
	g.GET("/saybye", examplepkg.SayBye())
	// End Example

	/* {
		"textToAnalysis":"loved it"
	} */
	g.POST("/sentiment", sentiment_controller.AnalyseTextSentimeter())

	//Examples for functions
	g.GET("/csv", data_controller.ReadCSV())

	/* {
		"colTarget":"ColunaIdade",
		"csv": arquivo.csv
	} */
	//Enviar arquivo csv e a coluna que deve ser usada pra gerar as visualizações
	g.POST("/grapher", data_controller.GenerateChartsFromGivenCsvAndTargetColumn())

	/* {
		"colTarget":"ColunaIdade",
		"csv": arquivo.csv
	} */
	//Enviar arquivo csv e a coluna que deve ser usada pra gerar as visualizações
	g.POST("/grapher&return", data_controller.GenerateChartsFromGivenCsvAndTargetColumnAndReturnThem())

	//Método pra recuperar os gráficos pela url
	g.POST("/plotter/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		//Todo := Aqui da problema por causa do inicio do root
		filePath := filepath.Join("data_pkg/output", filename)
		fmt.Println(filePath)

		// Check if the file exists
		_, err := os.Stat(filePath)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "File not found",
			})
			return
		}

		// Set the content type to image/png
		c.Header("Content-Type", "image/png")

		// Serve the file as-is
		c.File(filePath)
	})

}
