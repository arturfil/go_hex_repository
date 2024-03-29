package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/arturfil/yt_ecomm/helpers"
	"github.com/arturfil/yt_ecomm/service/auth"
	"github.com/arturfil/yt_ecomm/types"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRotues(router *chi.Mux) {
	router.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		res := struct {
			Msg  string
			Code int
		}{
			Msg:  "Health Check",
			Code: 200,
		}

		jsonStr, err := json.Marshal(res)
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonStr)
	})
	router.Post("/login", h.handleLogin)
	router.Post("/register", h.handleRegister)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

// handleRegister - will create a new user in the database
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// payload
	var payload types.RegisterUserPayload
	if err := helpers.ReadJSON(r, &payload); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate the payload
	if err := helpers.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		helpers.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

    fmt.Println("email -> ", payload.Email)

	// check is user exists
	user, err := h.store.GetUserByEmail(payload.Email)
    fmt.Println("USER HERE -> ", user)
	if err == nil {
		helpers.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email already %s exists", payload.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		helpers.WriteError(w, http.StatusInternalServerError, err)

	}

	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		helpers.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	helpers.WriteJSON(w, http.StatusCreated, nil)
}
