package main

import (
	"log"
	"net/http"
	"time"
	"wedding_api/internal/handler"
	"wedding_api/internal/repository"
	"wedding_api/internal/service"

	"github.com/gorilla/mux"
)



func main() {
	db, err := repository.NewDB()
	if err != nil {
		log.Fatal("Cannot connect to database.")
	}

	// Register all services
	dao := repository.NewDAO(db)
	userService := service.NewUserService(dao)

	// Register the handlers
	handlers := handler.MakeHandlers(&userService)

	// Start the HTTP server
	router := mux.NewRouter()
	router.HandleFunc("/api/user/{userId}", handlers.UserHandler)

	srv := &http.Server{
		Handler: router,
		Addr: ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())	
}