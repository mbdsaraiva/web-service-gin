package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = "secret" // Chave para assinar o token

// Estruturas para dados
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

type Vendor struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Dados em memória
var users = []Credentials{
	{Username: "admin", Password: "admin"},
}

var products = []Product{
	{ID: 1, Name: "Produto 1", Price: "10.00"},
	{ID: 2, Name: "Produto 2", Price: "20.00"},
}

var vendors = []Vendor{
	{ID: 1, Name: "Vendedor 1"},
	{ID: 2, Name: "Vendedor 2"},
}

// Middleware para autenticação JWT
func authenticateJWT(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		c.Abort()
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		c.Abort()
		return
	}

	c.Next()
}

// Funções para endpoints
func login(c *gin.Context) {
	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	// Verifica as credenciais
	for _, u := range users {
		if u.Username == creds.Username && u.Password == creds.Password {
			// Gera o token JWT
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": creds.Username,
				"exp":      time.Now().Add(time.Hour * 1).Unix(), // Expira em 1 hora
			})

			tokenString, err := token.SignedString([]byte(secretKey))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"token": tokenString})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid username or password"})
}

// Endpoints de produtos
func getProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func addProduct(c *gin.Context) {
	var product Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	product.ID = len(products) + 1
	products = append(products, product)
	c.JSON(http.StatusCreated, product)
}

// Endpoints de vendedores
func getVendors(c *gin.Context) {
	c.JSON(http.StatusOK, vendors)
}

func addVendor(c *gin.Context) {
	var vendor Vendor
	if err := c.BindJSON(&vendor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	vendor.ID = len(vendors) + 1
	vendors = append(vendors, vendor)
	c.JSON(http.StatusCreated, vendor)
}

// Endpoint protegido
func protectedEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "You are authorized!"})
}

func main() {
	r := gin.Default()

	// Endpoints públicos
	r.POST("/login", login)
	r.GET("/products", getProducts)
	r.GET("/vendors", getVendors)

	// Endpoints para adicionar (protegidos com JWT)
	r.POST("/products", authenticateJWT, addProduct)
	r.POST("/vendors", authenticateJWT, addVendor)

	// Endpoint protegido
	r.GET("/protected", authenticateJWT, protectedEndpoint)

	// Inicia o servidor
	r.Run(":8080")
}