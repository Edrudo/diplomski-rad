package services

import (
	"http3-server-poc/internal/domain/models"

	"github.com/pkg/errors"
)

type PartStoringService struct {
	partsRepository      PartsRepository
	processingEngineChan chan string
}

func NewPartStoringService(
	partsRepository PartsRepository,
	processingEngineChan chan string,
) *PartStoringService {
	return &PartStoringService{
		partsRepository:      partsRepository,
		processingEngineChan: processingEngineChan,
	}
}

// StorePart stores a part in the storage with the other relevant parts
func (i *PartStoringService) StorePart(part models.Part) error {
	errctx := func(err error) error {
		return errors.WithMessagef(err, "part storing service, store part")
	}

	// check if this part already exists in stoage
	parts, ok, err := i.partsRepository.GetPartsList(part.DataHash)
	if err != nil {
		return errctx(err)
	}

	// if it exists then do a noop
	if ok {
		for _, part := range parts {
			if part.PartNumber == part.PartNumber {
				return nil
			}
		}
	}

	// otherwise store the part
	err = i.partsRepository.StorePart(part)
	if err != nil {
		return errctx(err)
	}

	if len(parts)+1 == part.TotalParts {
		i.processingEngineChan <- part.DataHash
	}

	return nil
}
