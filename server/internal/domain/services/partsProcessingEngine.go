package services

import (
	"encoding/json"

	"http3-server-poc/cmd/api/config"
	"http3-server-poc/internal/domain/models"
)

type PartsProcessingEngine struct {
	dataHashChan      chan string
	partsRepository   PartsRepository
	geoDataRepository GeodataRepository
	imageStore        ImageStore
	jsonStore         JsonStore
}

func NewPartsProcessingEngine(
	dataHashChan chan string,
	partsRepository PartsRepository,
	geoDataRepository GeodataRepository,
	imageStore ImageStore,
	jsonStore JsonStore,
) *PartsProcessingEngine {
	return &PartsProcessingEngine{
		dataHashChan:      dataHashChan,
		partsRepository:   partsRepository,
		geoDataRepository: geoDataRepository,
		imageStore:        imageStore,
		jsonStore:         jsonStore,
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

	// unmarshall data to json
	var geoshot models.Geoshot
	err = json.Unmarshal(dataBytes, &geoshot)
	if err != nil {
		// add logging
		return
	}

	// store image to filesystem
	imagePath, err := e.imageStore.StoreImage(dataHash, geoshot.Image)

	// store json to filesystem
	jsonBytes, err := json.Marshal(geoshot)
	if err != nil {
		// add logging
	}
	jsonPath, err := e.jsonStore.StoreJson(dataHash, jsonBytes)

	if config.Cfg.DatabaseEnabled {
		// save geo metadata
		err = e.geoDataRepository.SaveGeoshot(geoshot, imagePath, jsonPath)
	}

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
