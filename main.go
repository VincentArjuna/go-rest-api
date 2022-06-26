package main

import (
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
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

func bookById(c *gin.Context){
	id:= c.Param("id")
	book, err := getBooksById(id)
	
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found."})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func checkoutBook(c *gin.Context){
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := getBooksById(id)
	
	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found."})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book Not Available."})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context){
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := getBooksById(id)
	
	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found."})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}


func getBooksById(id string)(*book, error){
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
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
	router.GET("/books/:id", bookById)
	//curl localhost:8080/books/2
	router.POST("/books", createBooks)
	//curl localhost:8080/books --include --header "Content-Type: application/json" -d @body.json --request "POST"
	router.PATCH("/checkout", checkoutBook)
	//curl localhost:8080/checkout?id=2 --request "PATCH"
	router.PATCH("/return", returnBook)
	//curl localhost:8080/return?id=2 --request "PATCH"
	router.Run("localhost:8080")
}