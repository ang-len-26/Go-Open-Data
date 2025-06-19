package controllers

import (
	"net/http"

	"github.com/ang-len-26/go-open-data-api/config"
	"github.com/ang-len-26/go-open-data-api/models"
	"github.com/gin-gonic/gin"
)

func GetLanguages(c *gin.Context) {
	query := `
		SELECT id, name, native_name, iso_code
		FROM languages
		ORDER BY name ASC
	`
	rows, err := config.DB.Query(c.Request.Context(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los idiomas", "details": err.Error()})
		return
	}
	defer rows.Close()

	var languages []models.Language
	for rows.Next() {
		var lang models.Language
		if err := rows.Scan(&lang.ID, &lang.Name, &lang.NativeName, &lang.ISOCode); err != nil {
			continue
		}
		languages = append(languages, lang)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  languages,
		"total": len(languages),
	})
}

func GetLanguageByISO(c *gin.Context) {
	iso := c.Param("iso")

	var lang models.Language
	query := `
		SELECT id, name, native_name, iso_code
		FROM languages
		WHERE iso_code = $1
	`

	err := config.DB.QueryRow(c.Request.Context(), query, iso).Scan(
		&lang.ID, &lang.Name, &lang.NativeName, &lang.ISOCode,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Idioma no encontrado"})
		return
	}

	c.JSON(http.StatusOK, lang)
}

func GetCountriesByLanguage(c *gin.Context) {
	iso := c.Param("iso")

	query := `
		SELECT c.id, c.name, c.capital, c.population, c.area
		FROM countries c
		JOIN country_languages cl ON c.id = cl.country_id
		JOIN languages l ON cl.language_id = l.id
		WHERE l.iso_code = $1
		ORDER BY c.name ASC
	`

	rows, err := config.DB.Query(c.Request.Context(), query, iso)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los países", "details": err.Error()})
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
