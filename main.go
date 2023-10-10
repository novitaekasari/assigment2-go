package main

import (
	"github.com/novitaekasari/assigment2-go/models"
	"github.com/novitaekasari/assigment2-go/routers"
)

func main() {
	models.ConnectDatabase()
	routers.StartServer().Run(":8080")
}