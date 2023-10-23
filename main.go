package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// fields must start with capital to be an exported field name aka "public field", which means it can be viewed by modules outside of this file
// `json:"id"` is used to represent the Json field name in the struct, and in Json it will convert the field name to lowercase
// vice versa, when the Json object is converted to the book struct the field will turn back to start with uppercase
// because we added the json bits at the back, our book can be serialized easily to JSON
type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

// handle GET all books
// this function takes a context (c), which stores all the info about the request such as query parameters, payload, headers
func getBooks(c *gin.Context) {
	// IndentedJSON will format the JSON for us
	// the HTTP status code will be OK, and the data is the books
	c.IndentedJSON(http.StatusOK, books)
}

// function to bind data in the request context to a book
func bookById(c *gin.Context) {
	// Param means that this is a path parameter like "/books/2" where 2 is the id
	// it's defined in the router path in Main below
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		// return a custom response saying 404 not found
		// gin.H is a shortcut to allow us to easily write custom JSON to be returned
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func checkoutBook(c *gin.Context) {
	// we check out books by ID
	// but this time we accept it as a query parameter instead of a path parameter so we can learn how it works
	// we need to add ok here because if the id doesn't exist then it is not "ok"
	id, ok := c.GetQuery("id")

	// if !ok is basically if ok == false
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	// check book quantity. We cannot let it be checked out if quantity is 0
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "No more books left"})
		return
	}

	// reduce the quantity of a type of book if it was checked out
	book.Quantity -= 1

	// return the checked out book
	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	// add the quantity of a type of book if it was returned
	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

// helper function to GET one book by ID
// we return a pointer to a book (*book) and a error because the book might not exist. This is shown by the (*book, error)
func getBookById(id string) (*book, error) {
	// loop through all books to look for the right book
	for i, b := range books {
		if b.ID == id {
			// we return &books[i] to get a pointer to the right book, so that we can modify the struct's fields from a different function
			// we return nil as an error if the right book is found
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func createBook(c *gin.Context) {
	// create a new variable of type book
	var newBook book

	// we need to use something from c to bind the JSON (which was part of the request payload) to the newBook variable
	// we're passing the pointer (&) to the newBook (which we can directly modify the fields of)
	// and we check whether we got an error (if the error is != null means we got an error) then we can directly return
	// rmb: Returning does not automatically return a response. The .BindJSON() method is what will handle sending the error response
	if err := c.BindJSON(&newBook); err != nil {
		// enter here if there is an error
		return
	}

	// move here if there are no erros, and now we bind the JSON to the newBook struct
	// which contains all the data that was returned to this endpoint which we can append to the books array (aka slice in Go)
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// setting up a gin router to direct http requests
func main() {
	router := gin.Default()

	// define a "localhost:8080/books" endpoint
	router.GET("/books", getBooks)
	router.GET("books/:id", bookById)
	router.POST("/books", createBook)
	router.PATCH("/return", returnBook)
	router.PATCH("/checkout", checkoutBook)
	router.Run("localhost:8080")
}
