package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/guiln/starter-log/logger"
	"github.com/guiln/starter-log/messages"
	"github.com/guiln/starter-user-api/app/config"
	v1 "github.com/guiln/starter-user-api/app/v1"
	"github.com/guiln/starter-user-api/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	lggr := buildLogger("None")
	cfg := config.NewAppConfig().WithEnvVarGetter()
	port := fmt.Sprintf(":%s", cfg.GetParameter("PORT"))

	router := buildRouter(lggr)
	{
		lggr.Info(messages.New(fmt.Sprintf("Starting at port: %s", port)).Message())

		if err := router.Run(port); err != nil {
			lggr.Error(messages.New("Error while starting router").WithError(err).Message())
		}
	}
}

func buildRouter(lggr logger.CompanyLogger) *gin.Engine {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/v1/user"

	v1UserHandler := v1.NewUserHandler(lggr)

	// v1User routes
	v1User := router.Group("/v1/user")
	{
		v1User.POST("/login", v1UserHandler.UserLogin)
		v1User.POST("/submit", nil)
	}

	// Swagger routes
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}

func buildLogger(correlationId string) logger.CompanyLogger {
	return logger.NewBuilder().
		WithCorrelationId(correlationId).
		Build()
}
