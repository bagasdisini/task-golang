package routes

import (
	"backend/handlers"
	"backend/pkg/middleware"
	"backend/pkg/mysql"
	"backend/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/users", middleware.Auth(h.ShowUsers)).Methods("GET")
	r.HandleFunc("/user/{id}", middleware.Auth(h.GetUserByID)).Methods("GET")
}
