package services

import "http3-server-poc/internal/domain/models"

type PartsRepository interface {
	DoesPartListExist(string) (bool, error)
	DeletePartList(string) error
	StorePart(models.Part) error
	GetNumberOfPartsInStorage(string) (int, error)
	GetPartsList(string) ([]models.Part, bool, error)
}
