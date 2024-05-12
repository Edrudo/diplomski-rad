package services

import "http3-server-poc/internal/domain/models"

type PartProcessingEngine struct {
	nameChan        chan string
	partsRepository PartsRepository
	partStore       DataStore
}

func NewPartProcessingEngine(
	nameChan chan string,
	repository PartsRepository,
	partStore DataStore,
) *PartProcessingEngine {
	return &PartProcessingEngine{
		nameChan:        nameChan,
		partsRepository: repository,
		partStore:       partStore,
	}
}

func (e *PartProcessingEngine) StartProcessing() {
	for {
		select {
		case partName := <-e.nameChan:
			go e.ProcessPart(partName)
		}
	}
}

func (e *PartProcessingEngine) ProcessPart(name string) {
	parts, ok, err := e.partsRepository.GetPartsList(name)
	if err != nil {
		// add logging
		return
	}
	if !ok {
		// add logging
		return
	}

	partsListLen := len(parts)
	partNumberPartMap := getPartsMapFromList(parts)

	partBytes := make([]byte, 0)
	for i := 1; i <= partsListLen; i++ {
		part, ok := partNumberPartMap[i]
		if !ok {
			// add logging
			return
		}

		partBytes = append(partBytes, part.PartData...)
	}

	err = e.partStore.StorePart(name, partBytes)

	// delete parts from memory
	err = e.partsRepository.DeletePartList(name)

	// add logging

	return
}

func getPartsMapFromList(parts []models.Part) map[int]models.Part {
	partNumberPartMap := make(map[int]models.Part)
	for _, part := range parts {
		partNumberPartMap[part.PartNumber] = part
	}

	return partNumberPartMap
}
