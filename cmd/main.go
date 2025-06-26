package main

import (
	"log"

	"github.com/ang-len-26/go-open-data-api/config"
	"github.com/ang-len-26/go-open-data-api/database"
	"github.com/ang-len-26/go-open-data-api/routes"
	"github.com/gin-contrib/cors"
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
	database.Connect()

	// Inicializar el router y registrar rutas
	r := gin.Default()

	// 👇 habilitar CORS para permitir peticiones desde el frontend
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://go-open-data.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	routes.RegisterCountryRoutes(r)
	routes.RegisterCityRoutes(r)
	routes.RegisterRegionRoutes(r)
	routes.RegisterLanguageRoutes(r)
	routes.RegisterCurrencyRoutes(r)

	// Iniciar el servidor
	port := config.GetEnv("PORT")
	if port == "" {
		port = "8080" // valor por defecto para desarrollo local
	}
	r.Run(":" + port)

}
