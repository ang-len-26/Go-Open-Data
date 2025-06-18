package routes

import (
	"github.com/ang-len-26/go-open-data-api/controllers"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes initializes the routes for the API
// and associates them with their respective handlers.
func RegisterCountryRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.GET("/countries", controllers.GetCountries)
		api.GET("/countries/:id", controllers.GetCountryByID)
		api.GET("/countries/:id/cities", controllers.GetCitiesByCountry)
		api.GET("/cities", controllers.GetCities)
		api.GET("/cities/:id", controllers.GetCityByID)
	}
}
