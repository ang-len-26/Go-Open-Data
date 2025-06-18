package controllers

import (
	"net/http"

	"github.com/ang-len-26/go-open-data-api/config"
	"github.com/ang-len-26/go-open-data-api/models"
	"github.com/gin-gonic/gin"
)

func GetSubregions(c *gin.Context) {
	rows, err := config.DB.Query(c.Request.Context(), `SELECT id, name FROM subregions ORDER BY name ASC`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener subregiones"})
		return
	}
	defer rows.Close()

	var subregions []models.Subregion
	for rows.Next() {
		var s models.Subregion
		if err := rows.Scan(&s.ID, &s.Name); err == nil {
			subregions = append(subregions, s)
		}
	}

	c.JSON(http.StatusOK, subregions)
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

	rows, err := config.DB.Query(c.Request.Context(), query, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener países de la subregión"})
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
