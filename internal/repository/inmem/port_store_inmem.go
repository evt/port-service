package inmem

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/evt/port-service/internal/domain"
)

type PortStore struct {
	data map[string]*Port
	mu   sync.RWMutex
}

func NewPortStore() *PortStore {
	return &PortStore{
		data: make(map[string]*Port),
	}
}

func (s *PortStore) GetPort(_ context.Context, id string) (*domain.Port, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	storePort, exists := s.data[id]
	if !exists {
		return nil, domain.ErrNotFound
	}

	domainPort, err := portStoreToDomain(storePort)
	if err != nil {
		return nil, fmt.Errorf("portStoreToDomain failed: %w", err)
	}

	return domainPort, nil
}

func (s *PortStore) CountPorts(_ context.Context) (int, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.data), nil
}

func (s *PortStore) CreateOrUpdatePort(ctx context.Context, p *domain.Port) error {
	if p == nil {
		return domain.ErrNil
	}

	storePort := portDomainToStore(p)

	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.data[storePort.ID]
	if exists {
		return s.updatePort(ctx, storePort)
	} else {
		return s.createPort(ctx, storePort)
	}
}

func (s *PortStore) createPort(_ context.Context, storePort *Port) error {
	if storePort == nil {
		return domain.ErrNil
	}

	// set created and updated at
	storePort.CreatedAt = time.Now()
	storePort.UpdatedAt = storePort.CreatedAt

	s.data[storePort.ID] = storePort

	return nil
}

func (s *PortStore) updatePort(_ context.Context, p *Port) error {
	if p == nil {
		return domain.ErrNil
	}

	// check if port exists
	storePort, exists := s.data[p.ID]
	if !exists {
		return domain.ErrNotFound
	}

	storePortCopy := storePort.Copy()

	storePortCopy.Name = p.Name
	storePortCopy.Code = p.Code
	storePortCopy.City = p.City
	storePortCopy.Country = p.Country
	storePortCopy.Alias = append([]string(nil), p.Alias...)
	storePortCopy.Regions = append([]string(nil), p.Regions...)
	storePortCopy.Coordinates = append([]float64(nil), p.Coordinates...)
	storePortCopy.Province = p.Province
	storePortCopy.Timezone = p.Timezone
	storePortCopy.Unlocs = append([]string(nil), p.Unlocs...)

	// set updated at
	storePortCopy.UpdatedAt = time.Now()

	s.data[p.ID] = storePortCopy

	return nil
}
