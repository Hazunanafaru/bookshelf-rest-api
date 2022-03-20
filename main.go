package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type book struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Year       int       `json:"year"`
	Author     string    `json:"author"`
	Summary    string    `json:"summary"`
	Publisher  string    `json:"publisher"`
	PageCount  int       `json:"pagecount"`
	ReadPage   int       `json:"readpage"`
	Reading    bool      `json:"reading"`
	Finished   bool      `json:"finished"`
	InsertedAt time.Time `json:"insertedat"`
	UpdatedAt  time.Time `json:"updatedat"`
}

var books = []book{}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookById)
	router.POST("/books", postBook)
	router.PUT("/books/:id", putBook)
	router.DELETE("/books/:id", deleteBook)

	router.Run("localhost:8080")
}
