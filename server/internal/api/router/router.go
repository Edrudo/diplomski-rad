package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	PartRoute = "/part"
)

type Controller interface {
	ProcessPart(requestContext *gin.Context)
}

// GenerateRoutingHandler creates a new routing handler which routes the requests accordingly
func GenerateRoutingHandler(
	controller Controller,
) (http.Handler, error) {
	gin.SetMode(gin.ReleaseMode)

	// Set up a fresh router with strict slash enabled.
	router := gin.New()
	router.RedirectTrailingSlash = true

	// Add middleware
	router.Use(gin.Logger())

	// IMAGE PROCESSING ENDPOINT
	{
		router.POST(PartRoute, controller.ProcessPart)
	}

	return router, nil
}
