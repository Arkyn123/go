package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test_service"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var user server.User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, statusCode, err := h.services.Users.CreateUser(user)
	if err != nil {
		c.AbortWithStatusJSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (h *Handler) FindAllUsers(c *gin.Context) {
	users, statusCode, err := h.services.Users.FindAllUsers()
	if err != nil {
		c.AbortWithStatusJSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *Handler) FindById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user, statusCode, err := h.services.FindById(id)
	if err != nil {
		c.AbortWithStatusJSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
