package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	VendorID    string  `json:"vendor_id"`
}

type vendor struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type user struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var products = []product{
	{ID: "1", Name: "Laptop", Description: "Laptop de 15 polegadas", Price: 1200.00, VendorID: "1"},
	{ID: "2", Name: "Smartphone", Description: "Smartphone 6GB RAM", Price: 600.00, VendorID: "2"},
}

var vendors = []vendor{
	{ID: "1", Name: "John's Electronics"},
	{ID: "2", Name: "TechWorld"},
}

var users = []user{
	{ID: "1", Username: "matheus", Password: "senha123"},
}

func main() {
	router := gin.Default()

	// Endpoints de Produtos
	router.GET("/products", getProducts)
	router.GET("/products/:id", getProductByID)
	router.POST("/products", postProduct)
	router.PUT("/products/:id", updateProduct)
	router.DELETE("/products/:id", deleteProduct)

	// Endpoints de Vendedores
	router.GET("/vendors", getVendors)
	router.GET("/vendors/:id", getVendorByID)
	router.POST("/vendors", postVendor)

	// Endpoints de Usuários
	router.GET("/users", getUsers)
	router.POST("/users", postUser)

	router.Run("localhost:8080")
}

func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

func getProductByID(c *gin.Context) {
	id := c.Param("id")
	for _, p := range products {
		if p.ID == id {
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

func postProduct(c *gin.Context) {
	var newProduct product
	if err := c.BindJSON(&newProduct); err != nil {
		return
	}
	products = append(products, newProduct)
	c.IndentedJSON(http.StatusCreated, newProduct)
}

func updateProduct(c *gin.Context) {
	id := c.Param("id")
	var updatedProduct product
	if err := c.BindJSON(&updatedProduct); err != nil {
		return
	}

	for i, p := range products {
		if p.ID == id {
			products[i] = updatedProduct
			c.IndentedJSON(http.StatusOK, updatedProduct)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

func deleteProduct(c *gin.Context) {
	id := c.Param("id")
	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "product deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

func getVendors(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, vendors)
}

func getVendorByID(c *gin.Context) {
	id := c.Param("id")
	for _, v := range vendors {
		if v.ID == id {
			c.IndentedJSON(http.StatusOK, v)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "vendor not found"})
}

func postVendor(c *gin.Context) {
	var newVendor vendor
	if err := c.BindJSON(&newVendor); err != nil {
		return
	}
	newVendor.ID = string(len(vendors) + 1)
	vendors = append(vendors, newVendor)
	c.IndentedJSON(http.StatusCreated, newVendor)
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func postUser(c *gin.Context) {
	var newUser user
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	newUser.ID = string(len(users) + 1)
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}
