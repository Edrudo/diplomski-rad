package controllermappers

import (
	"http3-server-poc/internal/api/controller"
	"http3-server-poc/internal/domain/models"
)

type ServerRequestMapper struct{}

func NewServerRequestMapper() *ServerRequestMapper {
	return &ServerRequestMapper{}
}

func (m *ServerRequestMapper) MapPartDtoToPartDomainModel(imagePart controller.Part) models.Part {
	return models.Part{
		DataHash:   imagePart.ImageHash,
		PartNumber: imagePart.PartNumber,
		TotalParts: imagePart.TotalParts,
		PartData:   imagePart.PartData,
	}
}
