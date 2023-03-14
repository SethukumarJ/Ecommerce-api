package http

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "ecommerce/cmd/api/docs"
	handler "ecommerce/pkg/api/handler"
	
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler) *ServerHTTP {
	engine := gin.New()

	// Use logger from Gin
	engine.Use(gin.Logger())
	
	// Swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Request JWT
	

	// Auth middleware
	// api := engine.Group("/api", middleware.AuthorizationMiddleware)

	engine.GET("/users", userHandler.FindByID)
	engine.POST("/api/users", userHandler.Save)
	engine.DELETE("users", userHandler.Delete)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
