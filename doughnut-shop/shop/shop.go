package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type order struct {
	ID       int `json:"id"`
	Quantity int `json:"quantity"`
	Total    int `json:"total"`
}

var orders = []order{
	{ID: 1, Quantity: 3, Total: 30},
	{ID: 2, Quantity: 4, Total: 40},
	{ID: 3, Quantity: 5, Total: 50},
}

func getOrders() []order {
	return orders
}

func getOrderById(orderId int) (order, error) {

	for _, o := range orders {
		if o.ID == orderId {
			return o, nil
		}
	}
	return order{}, errors.New("order with id not found")
}

func addOrder(newOrder order) (bool, error) {
	orders = append(orders, newOrder)
	return true, nil
}

func getOrdersHandle(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, getOrders())
}

func getOrdersByIdHandle(c *gin.Context) {
	var err error

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	foundOrder, err := getOrderById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	}

	c.IndentedJSON(http.StatusFound, foundOrder)
}

func postOrderHandle(c *gin.Context) {
	var newOrder order
	var err error
	//Call BindJson to bind the received JSON to newOrder
	err = c.BindJSON(&newOrder)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	// add the order to the list of orders
	_, err = addOrder(newOrder)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}

	c.IndentedJSON(http.StatusCreated, newOrder)
}

func main() {
	router := gin.Default()
	router.GET("/orders", getOrdersHandle)
	router.GET("/orders/:id", getOrdersByIdHandle)
	router.POST("/orders", postOrderHandle)

	router.Run()
}
