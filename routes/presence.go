package routes

import (
	"backend/handlers"
	"backend/pkg/middleware"
	"backend/pkg/mysql"
	"backend/repositories"

	"github.com/gorilla/mux"
)

func PresenceRoutes(r *mux.Router) {
	PresenceRepository := repositories.RepositoryPresence(mysql.DB)
	h := handlers.HandlerPresence(PresenceRepository)

	r.HandleFunc("/Presences", middleware.Auth(h.ShowPresences)).Methods("GET")
	r.HandleFunc("/Presence/{id}", middleware.Auth(h.GetPresenceByID)).Methods("GET")
	r.HandleFunc("/Presence", middleware.Auth(h.CreatePresence)).Methods("POST")
	r.HandleFunc("/Presence/{id}", middleware.Auth(h.UpdatePresence)).Methods("PATCH")
}
