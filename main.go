package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Estruturas para Produto e Vendedor
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

// Banco de dados em memória
var products = []product{
	{ID: "1", Name: "Laptop", Description: "Laptop de 15 polegadas", Price: 1200.00, VendorID: "1"},
	{ID: "2", Name: "Smartphone", Description: "Smartphone 6GB RAM", Price: 600.00, VendorID: "2"},
}

var vendors = []vendor{
	{ID: "1", Name: "John's Electronics"},
	{ID: "2", Name: "TechWorld"},
}

func main() {
	router := gin.Default()

	// Roteamento
	router.GET("/products", getProducts)
	router.GET("/products/:id", getProductByID)
	router.POST("/products", postProduct)
	router.PUT("/products/:id", updateProduct)
	router.DELETE("/products/:id", deleteProduct)

	router.GET("/vendors", getVendors)
	router.GET("/vendors/:id", getVendorByID)

	router.Run("localhost:8080")
}

// Funções de Produtos
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

// Funções de Vendedores
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