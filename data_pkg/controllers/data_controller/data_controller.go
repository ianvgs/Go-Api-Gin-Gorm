package data_controller

import (
	"archive/zip"
	"bytes"
	"fmt"
	"image/color"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// Todo := Aqui da problema quando muda de pasta pq tem que pegar do root
var outputFolder = "./data_pkg/output"

func ReadCSV() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(http.StatusForbidden, gin.H{
			"Message": "Blocked Route",
		})
		return

		//OS Open CSV
		csvFile, err := os.Open("data/inventory.csv")
		if err != nil {
			log.Fatal(err)
		}
		//Function to close archive on OS
		defer csvFile.Close()
		//Transform in dataframe
		df := dataframe.ReadCSV(csvFile)

		/* $$$FUNCTIONS$$$ */

		//1)Shape of dataset
		/*row, col := df.Dims() */
		/*fmt.Println("Shape of df, rows:", row, "columns:", col) */

		/*2)Tamanho das linhas */
		/*fmt.Println("Rows:", df.Nrow()) */

		/*3)Tamanho das colunas*/
		/*fmt.Println("Cols:", df.Ncol()) */

		/*4)Nomes das colunas */
		/*fmt.Println("Nomes:", df.Names()) */

		/*5)Tipos das colunas */
		/*fmt.Println("Tipos:", df.Types()) */

		//6)Criar tabela no terminal
		/*fmt.Println("Describe:", df.Describe()) */

		//7)Selecionar coluna por nome (case sensitive)
		/* fmt.Println("SELECT:", df.Select("Model")) */

		//8)Selecionar coluna por index
		/* fmt.Println("SELECT:", df.Select(0)) */

		//9)Selecionar varias colunas
		/*var colunas = []string{"Model", "Make"} */
		/* var colunas = []string
		colunas = append(colunas, "Model")
		colunas = append(colunas, "Make") */
		/* 		fmt.Println("SELECT:", df.Select(colunas)) */

		//10)Selecionar Rows por index
		/* fmt.Println("SELECT ROWS:", df.Subset(0)) */

		//Aplicando FUNÇÕES
		//ds = data series
		ds := df.Col("Year")

		fmt.Println(ds.IsNaN())

		//Mean()

		//Gera a media por função do dataframe
		dsMean := ds.Mean()
		fmt.Println(dsMean)

		//Gera a media utilizando a biblioteca stat
		statsMean := stat.Mean(ds.Float(), nil)
		fmt.Println(statsMean)

		// Use the Filter method
		filteredDF := df.Filter(dataframe.F{
			Colname:    "Make",
			Comparator: "==",
			Comparando: "Ford",
		})

		fmt.Println(filteredDF)

		c.JSON(http.StatusOK, gin.H{
			"Result": "Readed CSV",
		})
	}
}

/*
var r regression.regression
r.SetObserved("rating")
r.SetVar(0, "Sugars")

//Loop csv records adding the training data

fo i, record := range trainingData {

	//skip the header

	if i == 0 {
		continue
	}

	ratingVal, err := strconv.ParseFloat(record[0],64)
	if err != nil {
		log.Fatal(err)
	}
	sugarsVal, err := strconv.ParseFloat(record[2],64)
	if err != nil {
		log.Fatal(err)
	}

	//Add these point to the regression value
	r.Train(regression.DataPoint(ratingVal, []float64{sugarsVal}))



} */

func GenerateChartsFromGivenCsvAndTargetColumn() gin.HandlerFunc {
	return func(c *gin.Context) {

		/* $$$ 1) Recebe o csv, a coluna target, abre e lê o arquivo */
		//Obtem o valor colTarget do formulário enviado
		receivedColTarget := c.PostForm("colTarget")
		if receivedColTarget == "" {
			c.JSON(400, gin.H{"error": "The property colTarget was not informed in form."})
			return
		}
		fmt.Println("RECEBIDO:", receivedColTarget)

		//Procura pelo csv chamado csv no form
		csvFile, err := c.FormFile("csv")
		if err != nil {
			c.JSON(400, gin.H{"error": "CSV file not provided"})
			return
		}

		// Open the uploaded file
		srcFile, err := csvFile.Open()
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to open uploaded file"})
			return
		}
		defer srcFile.Close()

		//TODO CHECAR SE A COLUNA INFORMADA EXISTE NO CSV, se nao da pal

		//Le no dataframe o csv
		df := dataframe.ReadCSV(srcFile)
		//Levantar os Nomes da coluna pra devolver como dica se der erro
		columnNamesAvaiable := df.Names()

		//Checa se tem a coluna informada, se não houver vem um erro de unknown column (Err)
		ds := df.Col(receivedColTarget)
		if ds.Err != nil {
			c.JSON(400, gin.H{"Erro": "Não existe a coluna informada no csv inserido", "Colunas Encontradas": columnNamesAvaiable})
			return
		}

		//Checa se a coluna do eixo Y é do tipo int ou float
		dsType := ds.Type()
		/* if reflect.TypeOf(ds) == reflect.TypeOf(int(0)) */
		if dsType != "int" && dsType != "float" {
			c.JSON(400, gin.H{"Message": "A coluna informada precisa ser do tipo Int ou Float"})
			return
		}

		/* $$$ 2) Formata as entradas necessárias para gerar o gráfico */
		//Define a labal para os dados/coluna do eixo Y
		yColumnLabel := receivedColTarget
		//Assimilate Y Val as selected column above
		yVals := df.Col(yColumnLabel).Float()

		// Create an empty array of strings to save path to graphs
		var graphUrls []string

		for _, colName := range df.Names() {

			//Não cria grafico com eixo X igual a Y
			if colName == yColumnLabel {
				continue
			}

			//Checa se a coluna do eixo Y é do tipo int ou float, e não gera ela se for
			ds := df.Col(colName)
			dsType := ds.Type()
			if dsType != "int" && dsType != "float" {
				continue
			}

			pts := make(plotter.XYs, df.Nrow())

			/* for i, floatVal := range df.Col(colName).Float() {
				pts[i].X = floatVal
				pts[i].Y = yVals[i]
			}
			*/

			//Verifica se a coluna que vai ser analisada é numérica, se não for, skipa ela
			for i, floatVal := range df.Col(colName).Float() {
				if !math.IsNaN(floatVal) && !math.IsInf(floatVal, 0) {
					pts[i].X = floatVal
					pts[i].Y = yVals[i]
				} else {
					fmt.Printf("Error at index %d: Invalid floating-point value\n", i)
					continue
				}
			}

			p := plot.New()

			//Label X
			p.X.Label.Text = colName

			//LabelY
			p.Y.Label.Text = yColumnLabel

			p.Add(plotter.NewGrid())

			s, err := plotter.NewScatter(pts)
			if err != nil {
				log.Fatal("NewScatter Error")
			}

			s.GlyphStyle.Color = color.RGBA{R: 233, B: 0, A: 255}
			s.GlyphStyle.Radius = vg.Points(3)

			p.Add(s)

			newFileName := receivedColTarget + "versus" + colName + "_scatter.png"

			err = p.Save(4*vg.Inch, 4*vg.Inch, filepath.Join(outputFolder, newFileName))

			endpointString := os.Getenv("API_URL") + "/plotter/" + newFileName
			graphUrls = append(graphUrls, endpointString)

			if err != nil {
				log.Fatal(err)
			}

		}

		c.JSON(http.StatusOK, gin.H{
			"links": graphUrls,
		})

	}

}

func GenerateChartsFromGivenCsvAndTargetColumnAndReturnThem() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Create a map to store image data for each plot
		imageData := make(map[string][]byte)

		/* $$$ 1) Recebe o csv, a coluna target, abre e lê o arquivo */
		//Obtem o valor colTarget do formulário enviado
		receivedColTarget := c.PostForm("colTarget")
		if receivedColTarget == "" {
			c.JSON(400, gin.H{"error": "The property colTarget was not informed in form."})
			return
		}
		fmt.Println("RECEBIDO:", receivedColTarget)

		//Procura pelo csv chamado csv no form
		csvFile, err := c.FormFile("csv")
		if err != nil {
			c.JSON(400, gin.H{"error": "CSV file not provided"})
			return
		}

		// Open the uploaded file
		srcFile, err := csvFile.Open()
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to open uploaded file"})
			return
		}
		defer srcFile.Close()

		//TODO CHECAR SE A COLUNA INFORMADA EXISTE NO CSV, se nao da pal

		//Le no dataframe o csv
		df := dataframe.ReadCSV(srcFile)
		//Levantar os Nomes da coluna pra devolver como dica se der erro
		columnNamesAvaiable := df.Names()

		//Checa se tem a coluna informada, se não houver vem um erro de unknown column (Err)
		ds := df.Col(receivedColTarget)
		if ds.Err != nil {
			c.JSON(400, gin.H{"Erro": "Não existe a coluna informada no csv inserido", "Colunas Encontradas": columnNamesAvaiable})
			return
		}

		//Checa se a coluna do eixo Y é do tipo int ou float
		dsType := ds.Type()
		/* if reflect.TypeOf(ds) == reflect.TypeOf(int(0)) */
		if dsType != "int" && dsType != "float" {
			c.JSON(400, gin.H{"Message": "A coluna informada precisa ser do tipo Int ou Float"})
			return
		}

		/* $$$ 2) Formata as entradas necessárias para gerar o gráfico */
		//Define a labal para os dados/coluna do eixo Y
		yColumnLabel := receivedColTarget
		//Assimilate Y Val as selected column above
		yVals := df.Col(yColumnLabel).Float()

		for _, colName := range df.Names() {

			//Não cria grafico com eixo X igual a Y
			if colName == yColumnLabel {
				continue
			}

			//Checa se a coluna do eixo Y é do tipo int ou float, e não gera ela se for
			ds := df.Col(colName)
			dsType := ds.Type()
			if dsType != "int" && dsType != "float" {
				continue
			}

			pts := make(plotter.XYs, df.Nrow())

			/* for i, floatVal := range df.Col(colName).Float() {
				pts[i].X = floatVal
				pts[i].Y = yVals[i]
			}
			*/

			//Verifica se a coluna que vai ser analisada é numérica, se não for, skipa ela
			for i, floatVal := range df.Col(colName).Float() {
				if !math.IsNaN(floatVal) && !math.IsInf(floatVal, 0) {
					pts[i].X = floatVal
					pts[i].Y = yVals[i]
				} else {
					fmt.Printf("Error at index %d: Invalid floating-point value\n", i)
					continue
				}
			}

			p := plot.New()

			//Label X
			p.X.Label.Text = colName

			//LabelY
			p.Y.Label.Text = yColumnLabel

			p.Add(plotter.NewGrid())

			s, err := plotter.NewScatter(pts)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error ao criar um NewScatter"})
				return

			}

			s.GlyphStyle.Color = color.RGBA{R: 233, B: 0, A: 255}
			s.GlyphStyle.Radius = vg.Points(3)

			p.Add(s)

			// Save the plot to a PNG file.
			var buf bytes.Buffer
			q, err := p.WriterTo(4*vg.Inch, 4*vg.Inch, "png")
			if err != nil {
				panic(err)
			}
			if _, err := q.WriteTo(&buf); err != nil {
				panic(err)
			}

			imageData[colName+".png"] = buf.Bytes()

			/* // Set the appropriate HTTP headers for serving an image
			c.Writer.Header().Set("Content-Type", "image/png")
			c.Writer.Header().Set("Content-Disposition", "attachment; filename=image.png")

			// Write the image data from the buffer as the HTTP response
			c.Writer.WriteHeader(http.StatusOK)
			_, err = c.Writer.Write(buf.Bytes())
			if err != nil {
				fmt.Println("Error writing image response:", err)
			} */

		}

		// Create a zip buffer to hold all the image files
		var zipBuffer bytes.Buffer
		zipWriter := zip.NewWriter(&zipBuffer)

		// Iterate through the image data map and create zip entries
		for fileName, data := range imageData {
			imgFileWriter, err := zipWriter.Create(fileName)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating zip file"})
				return
			}

			// Write the image data to the zip file
			_, err = imgFileWriter.Write(data)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error writing image to zip file"})
				return
			}
		}

		// Close the zip archive
		if err := zipWriter.Close(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error closing zip archive"})
			return
		}

		// Get the size of the ZIP data
		contentLength := len(zipBuffer.Bytes())
		fmt.Println("LENGTH", len(zipBuffer.Bytes()))

		// Set the appropriate HTTP headers for serving a zip file
		c.Writer.Header().Set("Content-Type", "application/zip")
		c.Writer.Header().Set("Content-Disposition", "attachment; filename=charts.zip")
		c.Writer.Header().Set("Content-Length", strconv.Itoa(contentLength)) // Convert to string
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Write(zipBuffer.Bytes())
	}
}
