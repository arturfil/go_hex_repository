package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/arturfil/yt_ecomm/service/user"
	"github.com/go-chi/chi/v5"
)

type APIServer struct {
    addr string
    db *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
    return &APIServer{
        addr: addr,
        db: db,
    }
}

func (s *APIServer) Run() error {

    router := chi.NewRouter()
    router.Mount("/api/v1", router)

    useStore := user.NewStore(s.db)
    userHandler := user.NewHandler(useStore)
    userHandler.RegisterRotues(router)

    log.Println("listening on", s.addr)

    return http.ListenAndServe(s.addr, router)
}

