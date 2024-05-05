package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/arturfil/go_repository_hex/service/product"
	"github.com/arturfil/go_repository_hex/service/user"
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

    productStore := product.NewStore(s.db)
    productHandler := product.NewHandler(productStore)
    productHandler.RegisterRoutes(router)

    log.Println("listening on", s.addr)

    return http.ListenAndServe(s.addr, router)
}

