package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-rabbit-sample/api-producer/domain"
)

func (h *handler) addPerson(c *gin.Context) {
	var person *domain.Person

	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newPerson, err := h.personService.AddPerson(person)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, newPerson)
}
