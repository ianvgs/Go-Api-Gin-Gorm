package main

import (
	"goagain/initializers"
	"goagain/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.ConnectMongoDB()
}

func main() {
	routes.HandleRequests()
}
