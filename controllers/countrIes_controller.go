package controllers

import (
	"net/http"
	"strconv"

	"github.com/ang-len-26/go-open-data-api/config"
	"github.com/ang-len-26/go-open-data-api/models"
	"github.com/gin-gonic/gin"
)

func GetCountries(c *gin.Context) {
	// Obtener parámetros de consulta
	region := c.Query("region")
	subregion := c.Query("subregion")
	name := c.Query("name")
	limit := c.DefaultQuery("limit", "10")
	offset := c.DefaultQuery("offset", "0")

	// Construir la consulta dinámica
	query := "FROM countries WHERE 1=1"
	var args []interface{}
	argCount := 1

	if region != "" {
		query += " AND region ILIKE $" + strconv.Itoa(argCount)
		args = append(args, "%"+region+"%")
		argCount++
	}
	if subregion != "" {
		query += " AND subregion ILIKE $" + strconv.Itoa(argCount)
		args = append(args, "%"+subregion+"%")
		argCount++
	}
	if name != "" {
		query += " AND name ILIKE $" + strconv.Itoa(argCount)
		args = append(args, "%"+name+"%")
		argCount++
	}

	// 1️⃣ Total count
	countQuery := "SELECT COUNT(*) " + query
	var total int
	err := config.DB.QueryRow(c.Request.Context(), countQuery, args...).Scan(&total)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al contar resultados", "details": err.Error()})
		return
	}

	// 2️⃣ Paginated query
	paginatedQuery := "SELECT id, name, capital, region, subregion, population, area " + query +
		" ORDER BY name ASC LIMIT $" + strconv.Itoa(argCount)
	args = append(args, limit)
	argCount++

	paginatedQuery += " OFFSET $" + strconv.Itoa(argCount)
	args = append(args, offset)

	// Ejecutar la consulta
	rows, err := config.DB.Query(c.Request.Context(), paginatedQuery, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al consultar la base de datos"})
		return
	}
	defer rows.Close()

	var countries []models.Country
	for rows.Next() {
		var country models.Country
		if err := rows.Scan(
			&country.ID, &country.Name, &country.Capital, &country.RegionID,
			&country.SubregionID, &country.Population, &country.Area,
		); err != nil {
			continue
		}
		countries = append(countries, country)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  countries,
		"total": total,
	})
}
