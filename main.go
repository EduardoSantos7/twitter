package main

import (
	"log"
	"github.com/EduardoSantos7/twitter/handlers"
	"github.com/EduardoSantos7/twitter/db"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Unable to connect to the DB")
		return
	}
	handlers.Handlers()
}