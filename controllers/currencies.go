package controllers

import (
	"net/http"

	"github.com/ang-len-26/go-open-data-api/config"
	"github.com/ang-len-26/go-open-data-api/models"
	"github.com/gin-gonic/gin"
)

func GetCurrencies(c *gin.Context) {
	query := `
		SELECT id, name, code, symbol
		FROM currencies
		ORDER BY name ASC
	`

	rows, err := config.DB.Query(c.Request.Context(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las monedas", "details": err.Error()})
		return
	}
	defer rows.Close()

	var currencies []models.Currency
	for rows.Next() {
		var currency models.Currency
		if err := rows.Scan(&currency.ID, &currency.Name, &currency.Code, &currency.Symbol); err == nil {
			currencies = append(currencies, currency)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  currencies,
		"total": len(currencies),
	})
}

func GetCurrencyByCode(c *gin.Context) {
	code := c.Param("code")

	var currency models.Currency
	query := `
		SELECT id, name, code, symbol
		FROM currencies
		WHERE code = $1
	`

	err := config.DB.QueryRow(c.Request.Context(), query, code).Scan(
		&currency.ID, &currency.Name, &currency.Code, &currency.Symbol,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Moneda no encontrada"})
		return
	}

	c.JSON(http.StatusOK, currency)
}

func GetCountriesByCurrency(c *gin.Context) {
	code := c.Param("code")

	query := `
		SELECT c.id, c.name, c.capital, c.population, c.area
		FROM countries c
		JOIN country_currencies cc ON c.id = cc.country_id
		JOIN currencies cu ON cc.currency_id = cu.id
		WHERE cu.code = $1
		ORDER BY c.name ASC
	`

	rows, err := config.DB.Query(c.Request.Context(), query, code)
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
