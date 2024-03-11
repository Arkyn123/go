package handler

import (
	"flag"
	"fmt"
	"test_service/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {

	mode := flag.String("mode", "debug", "Application mode (debug, release, test)")
	flag.Parse()
	fmt.Println("GIN_MODE:", *mode)

	gin.SetMode(*mode)

	router := gin.Default()

	router.GET("test", h.Get)

	auth := router.Group("/auth")
	{
		auth.POST("/", h.Auth)
	}

	users := router.Group("/users")
	{
		users.POST("/", h.CreateUser)
		users.GET("/", h.FindAllUsers)
		users.GET("/:id", h.FindById)
	}

	return router
}
