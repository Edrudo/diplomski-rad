package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Controller struct {
	logger              *zap.Logger
	partsStoringService PartsStoringService
	serverRequestMapper ServerRequestMapper
}

func NewController(
	partsStoringService PartsStoringService,
	logger *zap.Logger,
	mapper ServerRequestMapper,
) *Controller {
	return &Controller{
		logger:              logger,
		partsStoringService: partsStoringService,
		serverRequestMapper: mapper,
	}
}

// ProcessPart processes the received part
func (c *Controller) ProcessPart(requestContext *gin.Context) {
	var httpStatusCode int

	// Read the whole body of the request
	body, err := io.ReadAll(requestContext.Request.Body)
	if err != nil {
		c.logger.Error(
			"processing part, error while reading request body",
		)
		httpStatusCode = http.StatusInternalServerError
		requestContext.Writer.WriteHeader(httpStatusCode)

		return
	}

	// Unmarshal the received body into a DTO struct
	var partDto Part

	err = json.Unmarshal(body, &partDto)
	if err != nil {
		c.logger.Warn(
			"processing part, error while unmarshalling",
		)
		httpStatusCode = http.StatusBadRequest
		requestContext.Writer.WriteHeader(httpStatusCode)

		return
	}

	// Map the DTO struct into a domain model
	part := c.serverRequestMapper.MapPartDtoToPartDomainModel(partDto)

	// store the part
	err = c.partsStoringService.StorePart(part)
	if err != nil {
		c.logger.Warn(
			fmt.Sprintf("processing part, failed to store part: %s", err.Error()),
		)
		httpStatusCode = http.StatusInternalServerError
		requestContext.Writer.WriteHeader(httpStatusCode)
	}
}
