package services

import (
	"context"

	"github.com/evt/port-service/internal/domain"
)

// PortRepository is a port repository for the port service
type PortRepository interface {
	CreateOrUpdatePort(ctx context.Context, port *domain.Port) error
	CountPorts(ctx context.Context) (int, error)
	GetPort(ctx context.Context, id string) (*domain.Port, error)
}

// PortService is a port service
type PortService struct {
	repo PortRepository
}

// NewPortService creates a new port service
func NewPortService(repo PortRepository) PortService {
	return PortService{
		repo: repo,
	}
}

// GetPort returns a port by id
func (s PortService) GetPort(ctx context.Context, id string) (*domain.Port, error) {
	return s.repo.GetPort(ctx, id)
}

// CountPorts returns the number of ports
func (s PortService) CountPorts(ctx context.Context) (int, error) {
	return s.repo.CountPorts(ctx)
}

// CreateOrUpdatePort creates or updates a port
func (s PortService) CreateOrUpdatePort(ctx context.Context, port *domain.Port) error {
	return s.repo.CreateOrUpdatePort(ctx, port)
}
