package main

import (
	"net/http"

	"github.com/OGARCIIA/examen-backend/domain"
	"github.com/OGARCIIA/examen-backend/infrastructure"
	"github.com/OGARCIIA/examen-backend/presentation"
	"github.com/gin-gonic/gin"
)

func main() {
	infrastructure.InitDatabase()
	infrastructure.DB.AutoMigrate(&domain.Product{}, &domain.Order{})

	r := gin.Default()

	r.Use(infrastructure.CORSMiddleware())

	r.GET("/generate-token", func(c *gin.Context) {
		token, err := infrastructure.GenerateJWT()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	api := r.Group("/api")
	api.Use(infrastructure.AuthMiddleware())
	{
		api.GET("/products", presentation.GetProducts)
		api.POST("/products", presentation.CreateProduct)
		api.PUT("/products/:id", presentation.UpdateProduct)
		api.DELETE("/products/:id", presentation.DeleteProduct)

		api.POST("/orders", presentation.CreateOrder)
	}

	r.Run(":8080")
}
