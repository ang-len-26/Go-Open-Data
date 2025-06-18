package controllers

import (
	"net/http"

	"github.com/ang-len-26/go-open-data-api/config"
	"github.com/ang-len-26/go-open-data-api/models"
	"github.com/gin-gonic/gin"
)

func GetRegions(c *gin.Context) {
	rows, err := config.DB.Query(c.Request.Context(), `SELECT id, name FROM regions ORDER BY name ASC`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener regiones"})
		return
	}
	defer rows.Close()

	var regions []models.Region
	for rows.Next() {
		var r models.Region
		if err := rows.Scan(&r.ID, &r.Name); err == nil {
			regions = append(regions, r)
		}
	}

	c.JSON(http.StatusOK, regions)
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

	rows, err := config.DB.Query(c.Request.Context(), query, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener países de la región"})
		return
	}
	defer rows.Close()

	var countries []models.Country
	for rows.Next() {
		var c models.Country
		if err := rows.Scan(&c.ID, &c.Name, &c.Capital, &c.Population, &c.Area); err == nil {
			countries = append(countries, c)
		}
	}

	c.JSON(http.StatusOK, countries)
}
