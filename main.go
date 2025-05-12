package main

import (
	"api-go/database"
    "api-go/handlers"
    "api-go/api"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialisation de la base de données
	database.InitDB()
	defer database.DB.Close()

	// Création du repository et des handlers
	personRepo := repositories.NewPersonRepository(database.DB)
	personHandler := handlers.NewPersonHandler(personRepo)

	// Configuration du routeur Gin
	router := gin.Default()

	// Routes API
	api := router.Group("/api/v1")
	{
		api.POST("/persons", personHandler.CreatePerson)
		api.GET("/persons", personHandler.GetAllPersons)
		api.GET("/persons/:id", personHandler.GetPerson)
		api.PUT("/persons/:id", personHandler.UpdatePerson)
		api.DELETE("/persons/:id", personHandler.DeletePerson)
	}

	// Démarrage du serveur
	router.Run(":8080")
}