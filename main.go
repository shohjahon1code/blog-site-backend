package main

import (
	"blog-backend/handlers"
	"blog-backend/middleware"
	"blog-backend/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	models.ConnectDatabase()

	r := mux.NewRouter()

	r.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/posts", middleware.AuthMiddleware(handlers.CreatePostHandler)).Methods("POST")
	r.HandleFunc("/posts", handlers.GetPostsHandler).Methods("GET")

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
