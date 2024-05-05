package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arturfil/go_repository_hex/types"
	"github.com/go-chi/chi/v5"
)

func TestUserServiceHandlers(t *testing.T) {
    userStore := &mockUserStore{}
    handler := NewHandler(userStore)

    

    t.Run("should return bad status if the user payload is invalid", func(t *testing.T) {
        payload := types.RegisterUserPayload{
            FirstName: "user",
            LastName: "las",
            Email: "invalidemail",
            Password: "12341234",
        }

        marshalled, _ := json.Marshal(payload)


        req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
        if err != nil {
            t.Fatal(err)
        }

        rr := httptest.NewRecorder()
        router := chi.NewRouter()

        router.HandleFunc("/register", handler.handleRegister)
        router.ServeHTTP(rr, req)

        if rr.Code != http.StatusBadRequest {
            t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
        }
    })

    t.Run("should create a user succesfully", func(t *testing.T) {
        payload := types.RegisterUserPayload{
            FirstName: "user",
            LastName: "las",
            Email: "valid2@gmail.com",
            Password: "12341234",
        }

        marshalled, _ := json.Marshal(payload)

        req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
        if err != nil {
            t.Fatal(err)
        }

        rr := httptest.NewRecorder()
        router := chi.NewRouter()

        router.HandleFunc("/register", handler.handleRegister)
        router.ServeHTTP(rr, req)

        if rr.Code != http.StatusCreated {
            t.Errorf("[register] expected status code %d, got %d, -> %v", http.StatusCreated, rr.Code, rr.Body)
        }
    })

}

type mockUserStore struct {}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
    return nil, nil
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
    return nil, nil
}

func (m *mockUserStore) CreateUser(user types.User) error {
    return nil
}



    
