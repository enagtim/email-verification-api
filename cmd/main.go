package main

import (
	"api/email-verification/configs"
	"api/email-verification/pkg/db"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8000",
		Handler: router,
	}
	fmt.Println("Server start on port 8000")
	server.ListenAndServe()
	defer db.Close()
}
