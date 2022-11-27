package controllers

import (
	"encoding/json"
	"fun-w-flags/database"
	"fun-w-flags/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PaginaInicial(c *gin.Context) {

}

func TodosPaises(w http.ResponseWriter, r *http.Request) {
	var p []models.Country_flag

	database.DB.Take(&p)
	json.NewEncoder(w).Encode(p)

}
