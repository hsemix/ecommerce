package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hsemix/ecommerce/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (r *APIServer) Run() error {
	router := mux.NewRouter()

	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(r.db)
	userService := user.NewHandler(userStore)
	userService.RegisterRoutes(subrouter)

	log.Println("Server Listening on ", r.addr)
	return http.ListenAndServe(r.addr, router)
}
