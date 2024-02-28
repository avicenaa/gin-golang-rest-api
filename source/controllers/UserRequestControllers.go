package controllers

import (
	"api-request/model"
	"api-request/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get all user request forms
// @Description Get all user request forms
// @ID get-all-user-request-forms
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Form "OK"
// @Router /api/view_all [get]
// @Tags API Request
func GetAllForms(c *gin.Context) {
	forms, err := services.GetAllRequestData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get all forms"})
		return
	}

	c.JSON(http.StatusOK, forms)
}

// @Summary Get user request form by ID
// @Description Get user request form by ID
// @ID get-user-request-form-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "Form ID"
// @Success 200 {object} model.Form "OK"
// @Router /api/view_detail/{id} [get]
// @Tags API Request
func GetFormByID(c *gin.Context) {
	formID := c.Param("id")

	form, err := services.GetRequestByID(formID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Form not found"})
		return
	}

	c.JSON(http.StatusOK, form)
}

// @Summary Create new user request form
// @Description Create new user request form
// @ID create-new-user-request-form
// @Accept  json
// @Produce  json
// @Param data body model.Form true "User Request Form Data"
// @Success 200 {string} string "OK"
// @Router /api/user/input [post]
// @Tags User API Request
func CreateForm(c *gin.Context) {
	var newForm model.Form
	if err := c.BindJSON(&newForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := services.AddRequestData(newForm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create new form"})
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Update user request form by ID
// @Description Update user request form by ID
// @ID update-user-request-form-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "Form ID"
// @Param data body model.Form true "Updated User Request Form Data"
// @Success 200 {string} string "OK"
// @Router /api/user/update/{id} [put]
// @Tags User API Request
func UpdateFormByID(c *gin.Context) {
	formID := c.Param("id")

	var updatedForm model.Form
	if err := c.BindJSON(&updatedForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := services.UpdateRequestData(formID, updatedForm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update form"})
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Delete user request form by ID
// @Description Delete user request form by ID
// @ID delete-user-request-form-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "Form ID"
// @Success 200 {string} string "OK"
// @Router /api/user/{id} [delete]
// @Tags User API Request
func DeleteFormByID(c *gin.Context) {
	formID := c.Param("id")

	if err := services.DeleteRequestData(formID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete form"})
		return
	}

	c.Status(http.StatusOK)
}
