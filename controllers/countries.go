package controllers

// queda pendiente agregar el controlador de países fronterizos
// ya que aun no se ha creado la tabla country_borders
// CREATE TABLE country_borders (
//     country_id INTEGER REFERENCES countries(id),
//     border_country_id INTEGER REFERENCES countries(id),
//     PRIMARY KEY (country_id, border_country_id)
// );

import (
	"net/http"
	"strconv"

	"github.com/ang-len-26/go-open-data-api/database"
	"github.com/ang-len-26/go-open-data-api/models"
	"github.com/gin-gonic/gin"
)

func GetCountryByID(c *gin.Context) {
	id := c.Param("id")
	countryID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var detail models.CountryDetail

	// Consulta principal
	query := `
	SELECT c.id, c.name, c.capital, c.population, c.area, r.name AS region, s.name AS subregion
	FROM countries c
	JOIN regions r ON c.region_id = r.id
	JOIN subregions s ON c.subregion_id = s.id
	WHERE c.id = $1
	`

	err = database.DB.QueryRow(c.Request.Context(), query, countryID).Scan(
		&detail.ID, &detail.Name, &detail.Capital, &detail.Population,
		&detail.Area, &detail.Region, &detail.Subregion,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "País no encontrado"})
		return
	}

	// Idiomas
	langQuery := `
	SELECT l.id, l.name, l.iso_code, l.native_name
	FROM languages l
	JOIN country_languages cl ON cl.language_id = l.id
	WHERE cl.country_id = $1
	`
	rows, _ := database.DB.Query(c.Request.Context(), langQuery, countryID)
	for rows.Next() {
		var lang models.Language
		rows.Scan(&lang.ID, &lang.Name, &lang.ISOCode, &lang.NativeName)
		detail.Languages = append(detail.Languages, lang)
	}
	rows.Close()

	// Monedas
	currQuery := `
	SELECT cu.id, cu.name, cu.code, cu.symbol
	FROM currencies cu
	JOIN country_currencies cc ON cc.currency_id = cu.id
	WHERE cc.country_id = $1
	`
	rows, _ = database.DB.Query(c.Request.Context(), currQuery, countryID)
	for rows.Next() {
		var cur models.Currency
		rows.Scan(&cur.ID, &cur.Name, &cur.Code, &cur.Symbol)
		detail.Currencies = append(detail.Currencies, cur)
	}
	rows.Close()

	// Ciudades
	cityQuery := `
	SELECT id, name, population, country_id, latitude, longitude
	FROM cities
	WHERE country_id = $1
	`
	rows, _ = database.DB.Query(c.Request.Context(), cityQuery, countryID)
	for rows.Next() {
		var city models.City
		rows.Scan(&city.ID, &city.Name, &city.Population, &city.CountryID, &city.Latitude, &city.Longitude)
		detail.Cities = append(detail.Cities, city)
	}
	rows.Close()

	c.JSON(http.StatusOK, detail)
}

func GetCitiesByCountry(c *gin.Context) {
	countryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	query := `
		SELECT id, name, population, latitude, longitude
		FROM cities
		WHERE country_id = $1
		ORDER BY population DESC
	`

	rows, err := database.DB.Query(c.Request.Context(), query, countryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al consultar ciudades"})
		return
	}
	defer rows.Close()

	var cities []models.City
	for rows.Next() {
		var city models.City
		if err := rows.Scan(&city.ID, &city.Name, &city.Population, &city.Latitude, &city.Longitude); err == nil {
			cities = append(cities, city)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  cities,
		"total": len(cities),
	})
}

func GetLanguagesByCountry(c *gin.Context) {
	countryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	query := `
		SELECT l.id, l.name, l.native_name, l.iso_code
		FROM country_languages cl
		JOIN languages l ON cl.language_id = l.id
		WHERE cl.country_id = $1
	`

	rows, err := database.DB.Query(c.Request.Context(), query, countryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al consultar idiomas"})
		return
	}
	defer rows.Close()

	var languages []models.Language
	for rows.Next() {
		var lang models.Language
		if err := rows.Scan(&lang.ID, &lang.Name, &lang.NativeName, &lang.ISOCode); err == nil {
			languages = append(languages, lang)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  languages,
		"total": len(languages),
	})
}

func GetCountryBorders(c *gin.Context) {
	countryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	query := `
		SELECT c.id, c.name, c.capital
		FROM country_borders cb
		JOIN countries c ON cb.border_country_id = c.id
		WHERE cb.country_id = $1
	`

	rows, err := database.DB.Query(c.Request.Context(), query, countryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al consultar países fronterizos"})
		return
	}
	defer rows.Close()

	var neighbors []models.Country
	for rows.Next() {
		var country models.Country
		if err := rows.Scan(&country.ID, &country.Name, &country.Capital); err == nil {
			neighbors = append(neighbors, country)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  neighbors,
		"total": len(neighbors),
	})
}
