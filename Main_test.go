package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetPersonEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/61612935ee236cb2ed8d4564", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetPersonEndpoint)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"_id":"61612935ee236cb2ed8d4564","Name":"ooo","email":"abc@123","password":"ZGVm47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU="}`

	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			strings.TrimSpace(rr.Body.String()), expected)
	} else {
		log.Println("GetPersonEndpoint - PASSED")
	}
}

func TestGetPostEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/posts/6160b337b478854bed753beb", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetPostEndpoint)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"_id":"6160b337b478854bed753beb","Caption":"ooo","ImageURL":"example.com","userid":"6160b15ab7ff0acfd7338bc0"}`

	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			strings.TrimSpace(rr.Body.String()), expected)
	} else {
		log.Println("GetPostEndpoint - PASSED")
	}
}

func TestGetAllPosts(t *testing.T) {
	req, err := http.NewRequest("GET", "/posts/users/6160b15ab7ff0acfd7338bc0?page=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllPosts)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `[{"Caption":"ooo","ImageURL":"example.com","_id":"6160b31bb478854bed753bea","userid":"6160b15ab7ff0acfd7338bc0"},{"Caption":"ooo","ImageURL":"example.com","_id":"6160b337b478854bed753beb","userid":"6160b15ab7ff0acfd7338bc0"}]`

	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			strings.TrimSpace(rr.Body.String()), expected)
	} else {
		log.Println("GetAllPosts - PASSED")
	}
}

func TestCreatePersonEndPoint(t *testing.T) {

	var jsonStr = []byte(`{    "Name":"testcase",
    "Email":"123",
    "Password":"xyz"}`)

	req, err := http.NewRequest("POST", "/users/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreatePersonEndpoint)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"success":"Upload successful"}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			strings.TrimSpace(rr.Body.String()), expected)
	} else {
		log.Println("CreatePersonEndpoint - PASSED")
	}
}

func TestCreatePostEndPoint(t *testing.T) {

	var jsonStr = []byte(`{      "Caption":"ooo",
    "ImageURL":"example.com",
    "UserID":"6160b15ab7ff0acfd7338bc0"`)

	req, err := http.NewRequest("POST", "/posts/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreatePostEndpoint)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"success":"Upload successful"}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			strings.TrimSpace(rr.Body.String()), expected)
	} else {
		log.Println("CreatePostEndpoint - PASSED")
	}
}
