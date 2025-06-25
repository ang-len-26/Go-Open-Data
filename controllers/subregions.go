package controllers

import (
	"net/http"

	"github.com/ang-len-26/go-open-data-api/database"
	"github.com/ang-len-26/go-open-data-api/models"
	"github.com/gin-gonic/gin"
)

func GetSubregions(c *gin.Context) {
	rows, err := database.DB.Query(c.Request.Context(), `
		SELECT id, name 
		FROM subregions 
		ORDER BY name ASC
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener subregiones"})
		return
	}
	defer rows.Close()

	var subregions []models.Subregion
	for rows.Next() {
		var subregion models.Subregion
		if err := rows.Scan(&subregion.ID, &subregion.Name); err == nil {
			subregions = append(subregions, subregion)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  subregions,
		"total": len(subregions),
	})
}

func GetCountriesBySubregion(c *gin.Context) {
	name := c.Param("name")

	query := `
		SELECT c.id, c.name, c.capital, c.population, c.area
		FROM countries c
		JOIN subregions s ON c.subregion_id = s.id
		WHERE LOWER(s.name) = LOWER($1)
		ORDER BY c.name ASC
	`

	rows, err := database.DB.Query(c.Request.Context(), query, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener países de la subregión"})
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
