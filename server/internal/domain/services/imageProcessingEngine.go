package services

import "http3-server-poc/internal/domain/models"

type ImageProcessingEngine struct {
	imageNameChan        chan string
	imagePartsRepository ImagePartsRepository
	imageStore           ImageStore
}

func NewImageProcessingEngine(
	imageNameChan chan string,
	imagePartsRepository ImagePartsRepository,
	imageStore ImageStore,
) *ImageProcessingEngine {
	return &ImageProcessingEngine{
		imageNameChan:        imageNameChan,
		imagePartsRepository: imagePartsRepository,
		imageStore:           imageStore,
	}
}

func (e *ImageProcessingEngine) StartProcessing() {
	for {
		select {
		case imageName := <-e.imageNameChan:
			go e.ProcessImage(imageName)
		}
	}
}

func (e *ImageProcessingEngine) ProcessImage(imageName string) {
	imageParts, ok, err := e.imagePartsRepository.GetImagePartsList(imageName)
	if err != nil {
		// add logging
		return
	}
	if !ok {
		// add logging
		return
	}

	imagePartsListLen := len(imageParts)
	partNumberImagePartMap := getImagePartsMapFromList(imageParts)

	imageBytes := make([]byte, 0)
	for i := 1; i <= imagePartsListLen; i++ {
		imagePart, ok := partNumberImagePartMap[i]
		if !ok {
			// add logging
			return
		}

		imageBytes = append(imageBytes, imagePart.PartData...)
	}

	err = e.imageStore.StoreImage(imageName, imageBytes)

	// delete image parts from memory
	err = e.imagePartsRepository.DeleteImagePartList(imageName)

	// add logging

	return
}

func getImagePartsMapFromList(imageParts []models.ImagePart) map[int]models.ImagePart {
	partNumberImagePartMap := make(map[int]models.ImagePart)
	for _, imagePart := range imageParts {
		partNumberImagePartMap[imagePart.PartNumber] = imagePart
	}

	return partNumberImagePartMap
}
