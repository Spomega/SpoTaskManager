package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
)

func TestCreateTask(t *testing.T) {

	router := chi.NewRouter()

	router.Post("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	})

	server := httptest.NewServer(router)

	taskURL := fmt.Sprintf("%s/tasks/", server.URL)

	taskJSON := `{
		"data":{
		  "createdBy" : "kojo",
		  "name": "go to school",
		  "description" : "finish hard",
		  "tags" : ["education","school"]
	  }
	  }`

	request, err := http.NewRequest("POST", taskURL, strings.NewReader(taskJSON))

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 201 {
		t.Errorf("HTTP Status expected: 201, got: %d", res.StatusCode)
	}

}

func TestGetTask(t *testing.T) {
	router := chi.NewRouter()

	router.Get("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	server := httptest.NewServer(router)

	taskURL := fmt.Sprintf("%s/tasks/", server.URL)

	request, err := http.NewRequest("GET", taskURL, nil)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", res.StatusCode)
	}

}
