package sentiment_controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grassmudhorses/vader-go/lexicon"
	"github.com/grassmudhorses/vader-go/sentitext"
)

type SentimentBody struct {
	TextToAnalysis string
}

func AnalyseTextSentimeter() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Declaro estrutura
		/* var body SentimentBody */
		//Faço o Bind

		//Bind lança erro e altera headers
		//c.BindJSON(&body)

		textToAnalysis := c.PostForm("textToAnalysis")

		//Lanço o erro se a propriedade nao existir
		/* if body.TextToAnalysis == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"warning": "Nao foi encontrado o parâmetro textToAnalysis para realizar essa tarefa.",
			})
			return
		} */

		fmt.Println("textToAnalysis", textToAnalysis)

		//Parseio o texto para analise
		parsedText := sentitext.Parse(textToAnalysis, lexicon.DefaultLexicon)

		fmt.Println("textToAnalysis", parsedText)
		//Gero Score de popularidade do texto parseado
		result := sentitext.PolarityScore(parsedText)

		c.JSON(http.StatusOK, gin.H{
			"Negative:": result.Negative,
			"Positive:": result.Positive,
			"Neutral:":  result.Neutral,
		})

	}
}
