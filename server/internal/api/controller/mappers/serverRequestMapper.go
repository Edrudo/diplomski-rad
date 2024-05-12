package controllermappers

import (
	"http3-server-poc/internal/api/controller"
	"http3-server-poc/internal/domain/models"
)

type ServerRequestMapper struct{}

func NewServerRequestMapper() *ServerRequestMapper {
	return &ServerRequestMapper{}
}

func (m *ServerRequestMapper) MapPartDtoToPartDomainModel(part controller.Part) models.Part {
	return models.Part{
		DataHash:   part.DataHash,
		PartNumber: part.PartNumber,
		TotalParts: part.TotalParts,
		PartData:   part.PartData,
	}
}
