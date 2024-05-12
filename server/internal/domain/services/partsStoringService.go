package services

import (
	"http3-server-poc/internal/domain/models"

	"github.com/pkg/errors"
)

type PartsStoringService struct {
	partsRepository           PartsRepository
	partsProcessingEngineChan chan string
}

func NewPartsStoringService(
	partsRepository PartsRepository,
	partsProcessingEngineChan chan string,
) *PartsStoringService {
	return &PartsStoringService{
		partsRepository:           partsRepository,
		partsProcessingEngineChan: partsProcessingEngineChan,
	}
}

func (i *PartsStoringService) StorePart(part models.Part) error {
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
		for _, p := range parts {
			if p.PartNumber == part.PartNumber {
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
		i.partsProcessingEngineChan <- part.DataHash
	}

	return nil
}
