package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type item struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
}

var items = []item{
	{ID: "1", Name: "Book", Quantity: 2},
	{ID: "2", Name: "Movie", Quantity: 4},
	{ID: "3", Name: "Table", Quantity: 5},
	{ID: "4", Name: "Bottle", Quantity: 3},
}

func getItems(context *gin.Context) {
    context.IndentedJSON(http.StatusOK, items)
}

func addItem(context *gin.Context) {
	var newItem item

	err := context.BindJSON(&newItem); 
	if err != nil {
		return
	}

	items = append(items, newItem)
	context.IndentedJSON(http.StatusCreated, newItem)
}

func getItemById(id string) (*item, error) {
	for i, t := range items {
		if t.ID == id {
			return &items[i], nil
		}
	}

	return nil, errors.New("item not found")
}

func getIndexById(id string) (int, error) {
	for i, t := range items {
		if t.ID == id {
			return i, nil
		}
	}

	return 0, errors.New("item not found")
}

func getItem(context *gin.Context) {
	id := context.Param("id")
	item, err := getItemById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, item)
}

func updateItem(context *gin.Context) {
	id := context.Param("id")
	foundItem, err := getItemById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item not found"})
		return
	}

	var body item

	context.Bind(&body)

	fmt.Print("found", foundItem)
	fmt.Print("body", body.Name)

	if len(body.Name) != 0 {
      foundItem.Name = body.Name
	}

	if body.Quantity != 0 {
      foundItem.Quantity = body.Quantity
	}

	context.IndentedJSON(http.StatusOK, foundItem)
}

func deleteItem(context *gin.Context) {
    id := context.Param("id")
	index, err := getIndexById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}

	items = append(items[:index], items[index+1:]...)
	context.IndentedJSON(http.StatusCreated, items)
}

func main() {
	router := gin.Default()
	router.GET("/items/:id", getItem)
	router.GET("/items", getItems)
	router.POST("/items", addItem)
	router.PATCH("/items/:id", updateItem)
	router.DELETE("/items/:id", deleteItem)
	router.Run()
}