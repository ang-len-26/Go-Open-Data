package controllers

import (
	"net/http"
	"strconv"

	"github.com/ang-len-26/go-open-data-api/config"
	"github.com/ang-len-26/go-open-data-api/models"
	"github.com/gin-gonic/gin"
)

func GetCities(c *gin.Context) {
	query := "FROM cities WHERE 1=1"
	args := []interface{}{}
	argCount := 1

	if name := c.Query("name"); name != "" {
		query += " AND name ILIKE $" + strconv.Itoa(argCount)
		args = append(args, "%"+name+"%")
		argCount++
	}
	if country := c.Query("country_id"); country != "" {
		query += " AND country_id = $" + strconv.Itoa(argCount)
		args = append(args, country)
		argCount++
	}
	if minPop := c.Query("min_population"); minPop != "" {
		query += " AND population >= $" + strconv.Itoa(argCount)
		args = append(args, minPop)
		argCount++
	}
	if maxPop := c.Query("max_population"); maxPop != "" {
		query += " AND population <= $" + strconv.Itoa(argCount)
		args = append(args, maxPop)
		argCount++
	}

	// Conteo total
	countQuery := "SELECT COUNT(*) " + query
	var total int
	err := config.DB.QueryRow(c.Request.Context(), countQuery, args...).Scan(&total)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al contar ciudades"})
		return
	}

	// Paginación
	limit := c.DefaultQuery("limit", "10")
	offset := c.DefaultQuery("offset", "0")
	query = "SELECT id, name, population, latitude, longitude, country_id " + query +
		" ORDER BY population DESC LIMIT $" + strconv.Itoa(argCount)
	args = append(args, limit)
	argCount++
	query += " OFFSET $" + strconv.Itoa(argCount)
	args = append(args, offset)

	rows, err := config.DB.Query(c.Request.Context(), query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al consultar ciudades"})
		return
	}
	defer rows.Close()

	var cities []models.City
	for rows.Next() {
		var city models.City
		if err := rows.Scan(
			&city.ID, &city.Name, &city.Population,
			&city.Latitude, &city.Longitude, &city.CountryID,
		); err == nil {
			cities = append(cities, city)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  cities,
		"total": total,
	})
}

func GetCityByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	query := `
		SELECT id, name, population, latitude, longitude, country_id
		FROM cities
		WHERE id = $1
	`

	var city models.City
	err = config.DB.QueryRow(c.Request.Context(), query, id).Scan(
		&city.ID, &city.Name, &city.Population, &city.Latitude, &city.Longitude, &city.CountryID,
	)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ciudad no encontrada"})
		return
	}

	c.JSON(http.StatusOK, city)
}
