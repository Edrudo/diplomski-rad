package services

import (
	"http3-server-poc/internal/domain/models"
)

type GeodataRepository interface {
	SaveGeoshot(
		geoshot models.Geoshot,
		imagePath string,
		jsonPath string,
	) error
}
