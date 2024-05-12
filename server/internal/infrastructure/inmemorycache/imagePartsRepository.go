package inmemorycache

import (
	"sync"

	"http3-server-poc/internal/domain/models"
)

type PartsRepository struct {
	mu               sync.Mutex
	dataHashPartsMap map[string][]models.Part
}

func NewPartsRepository() *PartsRepository {
	return &PartsRepository{
		dataHashPartsMap: map[string][]models.Part{},
	}
}

func (r *PartsRepository) DoesPartListExist(dataHash string) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, exists := r.dataHashPartsMap[dataHash]
	return exists, nil
}

func (r *PartsRepository) DeletePartList(dataHash string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.dataHashPartsMap, dataHash)
	return nil
}

func (r *PartsRepository) StorePart(part models.Part) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	parts, ok := r.dataHashPartsMap[part.DataHash]
	if !ok {
		parts = make([]models.Part, 0)
	}

	parts = append(parts, part)
	r.dataHashPartsMap[part.DataHash] = parts
	return nil
}

func (r *PartsRepository) GetNumberOfPartsInStorage(dataHash string) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	parts, ok := r.dataHashPartsMap[dataHash]
	if !ok {
		return 0, nil
	}

	return len(parts), nil
}

func (r *PartsRepository) GetPartsList(dataHash string) ([]models.Part, bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	parts, ok := r.dataHashPartsMap[dataHash]

	return parts, ok, nil
}
