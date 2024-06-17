package main

import (
	"store-be-golang/models"
	"store-be-golang/routers"
)

func main() {
	models.ConnectDatabase()
	
	var PORT = ":8080"
	routers.StartServer().Run(PORT)
}