package services

import "http3-server-poc/internal/domain/models"

type PartsRepository interface {
	DoesPartListExist(imageHash string) (bool, error)
	DeletePartList(imageHash string) error
	StorePart(imagePart models.Part) error
	GetNumberOfPartsInStorage(imageHash string) (int, error)
	GetPartsList(imageHash string) ([]models.Part, bool, error)
}
