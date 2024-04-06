package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/hsemix/ecommerce/types"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "semix.me@gmail.com",
			Password:  "test",
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(marshalled))

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Unexpected status code, expected %v got %v", http.StatusBadRequest, rr.Code)
		}
	})
}

type mockUserStore struct {
}

func (s *mockUserStore) GetUserByEmail(email string) (*types.User, error) {

	return nil, fmt.Errorf("user not found: %v", email)
}

func (s *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (s *mockUserStore) CreateUser(user types.User) error {
	return nil
}
