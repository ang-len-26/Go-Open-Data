package routes

import (
	"github.com/ang-len-26/go-open-data-api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRegionRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.GET("/regions", controllers.GetRegions)
		api.GET("/regions/:name/countries", controllers.GetCountriesByRegion)

		api.GET("/subregions", controllers.GetSubregions)
		api.GET("/subregions/:name/countries", controllers.GetCountriesBySubregion)
	}
}
