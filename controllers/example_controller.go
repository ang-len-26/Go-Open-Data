package controllers

import (
	"context"
	"net/http"

	"github.com/ang-len-26/go-open-data-api/config"
	"github.com/ang-len-26/go-open-data-api/models"
	"github.com/gin-gonic/gin"
)

func GetPublicData(c *gin.Context) {
	rows, err := config.DB.Query(context.Background(), "SELECT id, name, country, population, latitude, longitude FROM cities")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al consultar la base de datos"})
		return
	}
	defer rows.Close()

	var cities []models.City
	for rows.Next() {
		var city models.City
		if err := rows.Scan(&city.ID, &city.Name, &city.Country, &city.Population, &city.Latitude, &city.Longitude); err != nil {
			continue
		}
		cities = append(cities, city)
	}

	c.JSON(http.StatusOK, cities)
}
