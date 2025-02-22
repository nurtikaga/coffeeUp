package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) listCoffee(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) createCoffee(c *gin.Context) {

}

func (h *Handler) updateCoffee(c *gin.Context) {

}

func (h *Handler) deleteCoffee(c *gin.Context) {

}
