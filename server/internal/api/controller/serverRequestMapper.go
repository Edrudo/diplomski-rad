package controller

import (
	"http3-server-poc/internal/domain/models"
)

type ServerRequestMapper interface {
	MapPartDtoToPartDomainModel(imagePart Part) models.Part
}
