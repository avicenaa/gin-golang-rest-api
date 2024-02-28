package controllers

import (
	"api-request/model"
	"api-request/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Update request form by ID
// @Description Update request form by ID
// @ID update-user-request-form-by-id-crawler
// @Accept  json
// @Produce  json
// @Param id path string true "Form ID"
// @Param data body model.Form true "Updated User Request Form Data"
// @Success 200 {string} string "OK"
// @Router /api/crawler/update/{id} [put]
// @Tags Crawler API Request
func AddCrawlerData(c *gin.Context) {
	formID := c.Param("id")

	var updatedForm model.FormCrawler
	if err := c.BindJSON(&updatedForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := services.AddCrawlerData(formID, updatedForm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update form"})
		return
	}

	c.Status(http.StatusOK)
}
