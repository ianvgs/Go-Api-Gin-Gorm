package convolutional_neural_network_controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tensorflow/tensorflow/tensorflow/go/op"
)

func AnalyseImage() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"Result:": "Call ok",
		})

	}
}

func PreProp() gin.HandlerFunc {
	return func(c *gin.Context) {

		s := op.NewScope()

		graph, err := s.Finalize()
		if err != nil {
			panic(err)
		}

		fmt.Println("graph", graph)

		c.JSON(http.StatusOK, gin.H{
			"Result:": "Call ok",
		})

	}
}
