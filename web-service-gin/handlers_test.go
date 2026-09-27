// cd .\web-service-gin\
// go test -v

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetAlbums(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/albums", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	getAlbums(c)

	fmt.Println("GET /albums")
	fmt.Println("Status:", rec.Code)
	fmt.Println("Response:", rec.Body.String())

	// Проверяем HTTP status
	if rec.Code != http.StatusOK {
		t.Errorf("got %d, want %d", rec.Code, http.StatusOK)
	}

	// Декодируем JSON
	var got []album

	err = json.Unmarshal(rec.Body.Bytes(), &got)
	if err != nil {
		t.Fatalf("failed to decode JSON: %v", err)
	}

	// Проверяем количество альбомов
	if len(got) != len(albums) {
		t.Fatalf(
			"expected %d albums, got %d",
			len(albums),
			len(got),
		)
	}

	// Проверяем ID каждого альбома
	for i := range albums {
		if got[i].ID != albums[i].ID {
			t.Errorf(
				"got %s, want %s",
				got[i].ID,
				albums[i].ID,
			)
		}
	}
}

func TestPostAlbums(t *testing.T) {
	body := `{
		"id": "4",
		"title": "The Modern Sound of Betty Carter",
		"artist": "Betty Carter",
		"price": 49.99
	}`

	req, err := http.NewRequest(
		http.MethodPost,
		"/albums",
		strings.NewReader(body),
	)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	postAlbums(c)

	fmt.Println("POST /albums")
	fmt.Println("Status:", rec.Code)
	fmt.Println("Response:", rec.Body.String())

	// Проверяем status code
	if rec.Code != http.StatusCreated {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusCreated,
			rec.Code,
		)
	}

	// Проверяем, что JSON ответа корректный
	var got album

	err = json.Unmarshal(rec.Body.Bytes(), &got)
	if err != nil {
		t.Fatalf("failed to decode JSON: %v", err)
	}

	fmt.Printf("Added album: %+v\n", got)

	// Проверяем данные альбома
	if got.ID != "4" {
		t.Errorf("expected ID 4, got %s", got.ID)
	}

	if got.Title != "The Modern Sound of Betty Carter" {
		t.Errorf(
			"expected title 'The Modern Sound of Betty Carter', got %s",
			got.Title,
		)
	}

	if got.Artist != "Betty Carter" {
		t.Errorf(
			"expected artist 'Betty Carter', got %s",
			got.Artist,
		)
	}

	if got.Price != 49.99 {
		t.Errorf(
			"expected price 49.99, got %f",
			got.Price,
		)
	}
}

func TestGetAlbumByID(t *testing.T) {
	// Создаём GET-запрос
	req, err := http.NewRequest(
		http.MethodGet,
		"/albums/1",
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	// Создаём recorder
	rec := httptest.NewRecorder()

	// Создаём Gin context
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	// Передаём параметр :id
	c.Params = gin.Params{
		{
			Key:   "id",
			Value: "1",
		},
	}

	// Вызываем handler
	getAlbumByID(c)

	fmt.Println("GET /albums/1")
	fmt.Println("Status:", rec.Code)
	fmt.Println("Response:", rec.Body.String())

	// Проверяем status code
	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusOK,
			rec.Code,
		)
	}

	// Декодируем JSON в album
	var got album

	err = json.Unmarshal(rec.Body.Bytes(), &got)
	if err != nil {
		t.Fatalf("failed to decode JSON: %v", err)
	}

	// Проверяем данные
	if got.ID != "1" {
		t.Errorf("expected ID 1, got %s", got.ID)
	}

	if got.Title != "Blue Train" {
		t.Errorf(
			"expected title 'Blue Train', got %s",
			got.Title,
		)
	}

	if got.Artist != "John Coltrane" {
		t.Errorf(
			"expected artist 'John Coltrane', got %s",
			got.Artist,
		)
	}

	if got.Price != 56.99 {
		t.Errorf(
			"expected price 56.99, got %f",
			got.Price,
		)
	}
}

func TestGetAlbumByIDNotFound(t *testing.T) {
	// Создаём GET-запрос
	req, err := http.NewRequest(
		http.MethodGet,
		"/albums/999",
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	// Создаём recorder
	rec := httptest.NewRecorder()

	// Создаём Gin context
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	// Передаём несуществующий ID
	c.Params = gin.Params{
		{
			Key:   "id",
			Value: "999",
		},
	}

	// Вызываем handler
	getAlbumByID(c)

	fmt.Println("GET /albums/999")
	fmt.Println("Status:", rec.Code)
	fmt.Println("Response:", rec.Body.String())

	// Проверяем 404
	if rec.Code != http.StatusNotFound {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusNotFound,
			rec.Code,
		)
	}

	// Декодируем JSON
	var got map[string]string

	err = json.Unmarshal(rec.Body.Bytes(), &got)
	if err != nil {
		t.Fatalf("failed to decode JSON: %v", err)
	}

	// Проверяем сообщение
	if got["message"] != "album not found" {
		t.Errorf(
			"expected message 'album not found', got %s",
			got["message"],
		)
	}
}
