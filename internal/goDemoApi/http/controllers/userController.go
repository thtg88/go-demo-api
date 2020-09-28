package controllers

import (
	"goDemoApi/internal/goDemoApi/database"
	"goDemoApi/internal/goDemoApi/http/types"
	"goDemoApi/internal/goDemoApi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UsersShow returns the user from the given id
func UsersShow(c *gin.Context) {
	var user models.User

	result := database.Instance().Debug().Preload("Role").Find(&user, c.Param("id"))

	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"Resource": user})
		return
	}

	errors := make(map[string][]string)
	errors["NotFound"] = []string{"User not found."}

	c.JSON(http.StatusNotFound, &types.ErrorJSONResponse{
		Errors:  errors,
		Message: "There was a problem processing your request",
	})
}
