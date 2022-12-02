package routes

import (
	"backend/handlers"
	"backend/pkg/middleware"
	"backend/pkg/mysql"
	"backend/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	authRepository := repositories.RepositoryAuth(mysql.DB)
	h := handlers.HandlerAuth(authRepository)

	r.HandleFunc("/register", h.RegisterSV).Methods("POST")
	r.HandleFunc("/registerUser", middleware.Auth(h.RegisterUser)).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")
}
