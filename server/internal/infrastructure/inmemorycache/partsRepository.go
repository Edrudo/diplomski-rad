package inmemorycache

import (
	"sync"

	"http3-server-poc/internal/domain/models"
)

type PartsRepository struct {
	mu               sync.Mutex
	partNamePartsMap map[string][]models.Part
}

func NewPartsRepository() *PartsRepository {
	return &PartsRepository{
		partNamePartsMap: map[string][]models.Part{},
	}
}

func (r *PartsRepository) DoesPartListExist(partName string) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, exists := r.partNamePartsMap[partName]
	return exists, nil
}

func (r *PartsRepository) DeletePartList(partName string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.partNamePartsMap, partName)
	return nil
}

func (r *PartsRepository) StorePart(part models.Part) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	parts, ok := r.partNamePartsMap[part.DataHash]
	if !ok {
		parts = make([]models.Part, 0)
	}

	parts = append(parts, part)
	r.partNamePartsMap[part.DataHash] = parts
	return nil
}

func (r *PartsRepository) GetNumberOfPartsInStorage(partName string) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	parts, ok := r.partNamePartsMap[partName]
	if !ok {
		return 0, nil
	}

	return len(parts), nil
}

func (r *PartsRepository) GetPartsList(partName string) ([]models.Part, bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	parts, ok := r.partNamePartsMap[partName]

	return parts, ok, nil
}
