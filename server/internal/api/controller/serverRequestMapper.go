package controller

import (
	"http3-server-poc/internal/domain/models"
)

type ServerRequestMapper interface {
	MapPartToPartDomainModel(imagePart Part) models.Part
}
