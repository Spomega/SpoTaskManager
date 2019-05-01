package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
)

func TestUserRegistration(t *testing.T) {

	router := chi.NewRouter()

	router.Post("/users/register", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	})

	server := httptest.NewServer(router)

	defer server.Close()

	registerURL := fmt.Sprintf("%s/users/register", server.URL)

	userJSON := `{"firstname":"Rosmi","lastname":"Shiju","email":"rose@xyz.com"}`

	request, err := http.NewRequest("POST", registerURL, strings.NewReader(userJSON))

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 201 {
		t.Errorf("HTTP Status expected: 201, got: %d", res.StatusCode)
	}

}

func TestUserLogin(t *testing.T) {

	router := chi.NewRouter()

	router.Post("/users/login", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	server := httptest.NewServer(router)
	defer server.Close()

	loginURL := fmt.Sprintf("%s/users/login", server.URL)

	loginJSON := `{"email":"spo@xyz.com","password":"spomega"}`

	request, err := http.NewRequest("POST", loginURL, strings.NewReader(loginJSON))

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("HTTP Status expected: OK, got: %d", res.StatusCode)
	}

}
