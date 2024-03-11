package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"net/http"
)

func (h *Handler) Get(c *gin.Context) {
	resp, err := http.Get(viper.GetString("testing"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.Data(http.StatusOK, "application/json", body)
}

func (h *Handler) Auth(c *gin.Context) {
	var auth any

	if err := c.BindJSON(&auth); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	body, statusCode, err := h.services.Auth.Login(viper.GetString("auth_url"), "application/json", auth)
	if err != nil {
		c.AbortWithStatusJSON(statusCode, err.Error())
		return
	}

	c.Data(statusCode, "application/json", body)
}
