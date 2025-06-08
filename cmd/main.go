package main

// @title           Open Data API
// @version         1.0
// @description     API pública de datos de países y ciudades.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Angel Lengua
// @contact.email  tuemail@ejemplo.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /api/v1

import (
	"log"

	"github.com/ang-len-26/go-open-data-api/config"
	"github.com/ang-len-26/go-open-data-api/database"
	_ "github.com/ang-len-26/go-open-data-api/docs" // Importa la documentación generada por Swagger
	"github.com/ang-len-26/go-open-data-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"     // Importa el paquete de archivos de Swagger para Gin
	ginSwagger "github.com/swaggo/gin-swagger" // Importa el paquete de Swagger para Gin
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

	// Inicializar el router y registrar rutas
	r := gin.Default()
	routes.RegisterRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Iniciar el servidor
	r.Run(":" + config.GetEnv("PORT"))

}
