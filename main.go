package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	//"errors"
)

type book struct{
	ID			string	`json:"id"`
	Title		string	`json:"title"`
	Author		string	`json:"author"`
	Quantity	int		`json:"quantity"`
}

var books = []book{
	{ID:"1", Title: "Book1", Author: "Me", Quantity: 2},
	{ID:"2", Title: "Book2", Author: "Me2", Quantity: 5},
	{ID:"3", Title: "Book3", Author: "Me3", Quantity: 6},
}

func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK, books)
}

func createBooks(c *gin.Context){
	var newBook book

	if err:= c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)

}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	//curl localhost:8080/books
	router.POST("/books", createBooks)
	//curl localhost:8080/books --include --header "Content-Type: application/json" -d @body.json --request "POST"
	router.Run("localhost:8080")
}