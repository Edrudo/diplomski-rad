package services

import "http3-server-poc/internal/domain/models"

type PartsProcessingEngine struct {
	dataHashChan    chan string
	partsRepository PartsRepository
	imageStore      ImageStore
}

func NewPartsProcessingEngine(
	dataHashChan chan string,
	partsRepository PartsRepository,
	imageStore ImageStore,
) *PartsProcessingEngine {
	return &PartsProcessingEngine{
		dataHashChan:    dataHashChan,
		partsRepository: partsRepository,
		imageStore:      imageStore,
	}
}

func (e *PartsProcessingEngine) StartProcessing() {
	for {
		select {
		case dataHash := <-e.dataHashChan:
			go e.ProcessParts(dataHash)
		}
	}
}

func (e *PartsProcessingEngine) ProcessParts(dataHash string) {
	parts, ok, err := e.partsRepository.GetPartsList(dataHash)
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

	dataBytes := make([]byte, 0)
	for i := 1; i <= partsListLen; i++ {
		part, ok := partNumberPartMap[i]
		if !ok {
			// add logging
			return
		}

		dataBytes = append(dataBytes, part.PartData...)
	}

	err = e.imageStore.StoreImage(dataHash, dataBytes)

	// delete parts from memory
	err = e.partsRepository.DeletePartList(dataHash)

	// TODO add logging

	return
}

func getPartsMapFromList(parts []models.Part) map[int]models.Part {
	partNumberPartMap := make(map[int]models.Part)
	for _, part := range parts {
		partNumberPartMap[part.PartNumber] = part
	}

	return partNumberPartMap
}
