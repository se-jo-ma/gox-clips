package routes

import (
	"github.com/se-jo-ma/gox-clips/gox-clips/api/handlers"
	"github.com/se-jo-ma/gox-clips/gox-clips/internal/services"
	"net/http"
)

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()
	authService := services.AuthService{}
	authHandler := handlers.AuthHandler{Service: authService}
	userService := services.UserService{}
	userHandler := handlers.UserHandler{Service: userService}
	fs := http.FileServer(http.Dir("internal/static"))

	router.Handle("/static/", http.StripPrefix("/static/", fs))
	router.HandleFunc("/user", userHandler.GetUser)
	router.HandleFunc("/", handlers.HomeHandler)
	router.HandleFunc("/load-content", handlers.LoadContentHandler)
	router.HandleFunc("/login", authHandler.Login)
	router.HandleFunc("/protected", authHandler.Protected)
	router.HandleFunc("/logout", authHandler.Logout)

	return router
}
