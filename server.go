package main

import (
	"fmt"
	_ "github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "gorm.io/gorm"
	"log"
	"net/http"
	"time"
	_"encoding/json"
	"task-assign/router"
	"task-assign/utils"
)

var env map[string]string


func main() {
	env, err := godotenv.Read()
	if err != nil {
		panic("Problem reading the .env")
	}
	utils.Connect()
	router.SetUpRoutes()
	
	port := env["PORT"]
	host := env["HOST"]

	fmt.Println("Listening at " + host + ":" + port + "...")
	server := &http.Server{
		Handler:  router.Rt,
		Addr:         host + ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
