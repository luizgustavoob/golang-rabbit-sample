package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-rabbit-sample/api-producer/domain"
)

type handler struct {
	personService domain.PersonService
}

func NewHandler(personService domain.PersonService) http.Handler {
	handler := &handler{personService}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Logger(), handler.recovery())
	router.POST("/people", handler.addPerson)

	return router
}

func (h *handler) recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if recovered := recover(); recovered != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
