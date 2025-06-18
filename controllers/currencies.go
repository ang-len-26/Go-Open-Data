package controllers

import (
	"net/http"

	"github.com/ang-len-26/go-open-data-api/config"
	"github.com/ang-len-26/go-open-data-api/models"
	"github.com/gin-gonic/gin"
)

func GetCurrencies(c *gin.Context) {
	rows, err := config.DB.Query(c, "SELECT id, name, code, symbol FROM currencies ORDER BY name ASC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las monedas"})
		return
	}
	defer rows.Close()

	var currencies []models.Currency
	for rows.Next() {
		var currency models.Currency
		if err := rows.Scan(&currency.ID, &currency.Name, &currency.Code, &currency.Symbol); err != nil {
			continue
		}
		currencies = append(currencies, currency)
	}

	c.JSON(http.StatusOK, currencies)
}

func GetCurrencyByCode(c *gin.Context) {
	code := c.Param("code")

	var currency models.Currency
	err := config.DB.QueryRow(c, `
		SELECT id, name, code, symbol
		FROM currencies
		WHERE code = $1
	`, code).Scan(&currency.ID, &currency.Name, &currency.Code, &currency.Symbol)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Moneda no encontrada"})
		return
	}

	c.JSON(http.StatusOK, currency)
}

func GetCountriesByCurrency(c *gin.Context) {
	code := c.Param("code")

	rows, err := config.DB.Query(c, `
		SELECT countries.id, countries.name, countries.capital, countries.population, countries.area
		FROM countries
		JOIN country_currencies ON countries.id = country_currencies.country_id
		JOIN currencies ON country_currencies.currency_id = currencies.id
		WHERE currencies.code = $1
		ORDER BY countries.name ASC
	`, code)

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
