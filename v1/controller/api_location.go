/*
 * Weather Aggregator API
 *
 * Aggregator API for collecting weather info of a location.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteLocationById - Delete location
func DeleteLocationById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// InsertLocation - Add location
func InsertLocation(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateLocationById - Update location
func UpdateLocationById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
