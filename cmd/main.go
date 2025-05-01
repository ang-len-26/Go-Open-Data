package main

import (
	"log"

	"github.com/ang-len-26/go-open-data-api/config"
	"github.com/ang-len-26/go-open-data-api/database"
	"github.com/ang-len-26/go-open-data-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Carga el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	// Inicializa la conexión a la base de datos
	config.InitDB()
	database.Connect()

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run(":" + config.GetEnv("PORT"))
	r.Run(":8080")
}
