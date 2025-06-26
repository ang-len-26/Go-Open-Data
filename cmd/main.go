package main

import (
	"log"
	"os"

	"github.com/ang-len-26/go-open-data-api/config"
	"github.com/ang-len-26/go-open-data-api/database"
	"github.com/ang-len-26/go-open-data-api/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Solo carga .env si estás en desarrollo local
	if os.Getenv("RENDER") == "" {
		_ = godotenv.Load()
	}

	// Conexión a DB
	database.Connect()

	// Router
	r := gin.Default()

	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://go-open-data.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Rutas
	routes.RegisterCountryRoutes(r)
	routes.RegisterCityRoutes(r)
	routes.RegisterRegionRoutes(r)
	routes.RegisterLanguageRoutes(r)
	routes.RegisterCurrencyRoutes(r)

	// Puerto (Render define automáticamente la variable PORT)
	port := config.GetEnv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Servidor escuchando en el puerto %s", port)
	r.Run(":" + port)
}
