package services

import (
	"context"

	"github.com/evt/port-service/internal/domain"
	"github.com/google/uuid"
)

// PortService is a port service
type PortService struct {
}

// NewPortService creates a new port service
func NewPortService() PortService {
	return PortService{}
}

// GetPort returns a port by id
func (s PortService) GetPort(ctx context.Context, id string) (*domain.Port, error) {
	randomID := uuid.New().String()
	return domain.NewPort(randomID, randomID, randomID, randomID, randomID,
		[]string{randomID}, []string{randomID}, []float64{1.0, 2.0}, randomID, randomID, nil)
}
