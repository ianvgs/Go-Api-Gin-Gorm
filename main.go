package main

import (
	"goagain/initializers"
	"goagain/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	routes.HandleRequests()
}
