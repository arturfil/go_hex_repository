package product

import (
	"net/http"

	"github.com/arturfil/go_repository_hex/helpers"
	"github.com/arturfil/go_repository_hex/types"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
    store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
    return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *chi.Mux) {
    router.Get("/products", h.handleGetProducts)
}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
    products, err := h.store.GetProducts()
    if err != nil {
        helpers.WriteError(w, http.StatusInternalServerError, err)
        return 
    }

    helpers.WriteJSON(w, http.StatusOK, products)
}
