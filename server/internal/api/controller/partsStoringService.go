package controller

import "http3-server-poc/internal/domain/models"

type PartsStoringService interface {
	StorePart(models.Part) error
}
