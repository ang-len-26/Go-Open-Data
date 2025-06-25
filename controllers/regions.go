package controllers

import (
	"net/http"

	"github.com/ang-len-26/go-open-data-api/database"
	"github.com/ang-len-26/go-open-data-api/models"
	"github.com/gin-gonic/gin"
)

func GetRegions(c *gin.Context) {
	rows, err := database.DB.Query(c.Request.Context(), `
		SELECT id, name 
		FROM regions 
		ORDER BY name ASC
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener regiones"})
		return
	}
	defer rows.Close()

	var regions []models.Region
	for rows.Next() {
		var region models.Region
		if err := rows.Scan(&region.ID, &region.Name); err == nil {
			regions = append(regions, region)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  regions,
		"total": len(regions),
	})
}

func GetCountriesByRegion(c *gin.Context) {
	name := c.Param("name")

	query := `
		SELECT c.id, c.name, c.capital, c.population, c.area
		FROM countries c
		JOIN regions r ON c.region_id = r.id
		WHERE LOWER(r.name) = LOWER($1)
		ORDER BY c.name ASC
	`

	rows, err := database.DB.Query(c.Request.Context(), query, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener países de la región"})
		return
	}
	defer rows.Close()

	var countries []models.Country
	for rows.Next() {
		var country models.Country
		if err := rows.Scan(&country.ID, &country.Name, &country.Capital, &country.Population, &country.Area); err == nil {
			countries = append(countries, country)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  countries,
		"total": len(countries),
	})
}
