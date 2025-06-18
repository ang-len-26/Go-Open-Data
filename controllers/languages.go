package controllers

import (
	"net/http"

	"github.com/ang-len-26/go-open-data-api/config"
	"github.com/ang-len-26/go-open-data-api/models"
	"github.com/gin-gonic/gin"
)

func GetLanguages(c *gin.Context) {
	rows, err := config.DB.Query(c, "SELECT id, name, native_name, iso_code FROM languages ORDER BY name ASC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los idiomas"})
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

	c.JSON(http.StatusOK, languages)
}

func GetLanguageByISO(c *gin.Context) {
	iso := c.Param("iso")

	var lang models.Language
	err := config.DB.QueryRow(c, `
		SELECT id, name, native_name, iso_code
		FROM languages
		WHERE iso_code = $1
	`, iso).Scan(&lang.ID, &lang.Name, &lang.NativeName, &lang.ISOCode)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Idioma no encontrado"})
		return
	}

	c.JSON(http.StatusOK, lang)
}

func GetCountriesByLanguage(c *gin.Context) {
	iso := c.Param("iso")

	rows, err := config.DB.Query(c, `
		SELECT countries.id, countries.name, countries.capital, countries.population, countries.area
		FROM countries
		JOIN country_languages ON countries.id = country_languages.country_id
		JOIN languages ON country_languages.language_id = languages.id
		WHERE languages.iso_code = $1
		ORDER BY countries.name ASC
	`, iso)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los países"})
		return
	}
	defer rows.Close()

	var countries []models.Country
	for rows.Next() {
		var country models.Country
		if err := rows.Scan(&country.ID, &country.Name, &country.Capital, &country.Population, &country.Area); err != nil {
			continue
		}
		countries = append(countries, country)
	}

	c.JSON(http.StatusOK, countries)
}
