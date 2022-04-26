package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type book struct {
	Id         int       `json:"id"`
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

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func getBookById(c *gin.Context) {
	id := c.Param("id")
	index, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
	}

	for _, b := range books {
		if b.Id == index {
			c.JSON(http.StatusOK, b)
			return
		}
	}
}

func postBook(c *gin.Context) {
	var newBook book
	id := len(books) + 1
	insertedDate := time.Now()
	updatedDate := insertedDate

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	pNewBook := &newBook
	if newBook.Author == "" {
		c.JSON(http.StatusBadRequest, "Book author didn't found")
	}

	pNewBook.Id = id
	pNewBook.InsertedAt = insertedDate
	pNewBook.UpdatedAt = updatedDate
	books = append(books, newBook)
	c.JSON(http.StatusCreated, id)
}

func putBook(c *gin.Context) {
	var updateBook book

	if err := c.BindJSON(&updateBook); err != nil {
		return
	}

	id := c.Param("id")
	index, _ := strconv.Atoi(id)
	pUpdateBook := &updateBook
	pUpdateBook.UpdatedAt = time.Now()

	for _, b := range books {
		if b.Id == index {
			books = append(books[:index], *pUpdateBook)
			c.JSON(http.StatusCreated, b)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "book Id not found"})
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")
	index, _ := strconv.Atoi(id)

	for _, b := range books {
		if b.Id == index {
			books = append(books[:index], books[index+1:]...)
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "book Id not found"})
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookById)
	router.POST("/books", postBook)
	router.PUT("/books/:id", putBook)
	router.DELETE("/books/:id", deleteBook)

	router.Run("localhost:8080")
}
