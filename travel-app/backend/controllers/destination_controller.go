package controllers

import (
	"net/http"

	"github.com/WesleySDev/travel-app/backend/database"
	model "github.com/WesleySDev/travel-app/backend/models"
	"github.com/gin-gonic/gin"
)

/*função para criar destino.:
Recebe JSON do frontend
Converte para struct Destination
Salva no banco PostgreSQL
Retorna o destino criado
*/

func CreateDestination(c *gin.Context) {
	var destination model.Destination

	if err := c.ShouldBindJSON(&destination); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "dados inválidos",
		})
		return
	}
	if err := database.DB.Create(&destination).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "falha ao criar destino",
		})
		return
	}
	c.JSON(http.StatusOK, destination)
}

/*
função para listar destinos:
Busca todos os destinos no banco PostgreSQL
Retorna a lista de destinos para o frontend
*/
func GetDestinations(c *gin.Context) {
	var destinations []model.Destination

	if err := database.DB.Find(&destinations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "falha ao buscar destinos",
		})
		return
	}
	c.JSON(http.StatusOK, destinations)
}

func UpdateDestination(c *gin.Context) {
	//=========
}
