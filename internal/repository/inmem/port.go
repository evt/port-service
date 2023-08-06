package inmem

import (
	"errors"
	"time"

	"github.com/evt/port-service/internal/domain"
)

type Port struct {
	ID          string
	Name        string
	Code        string
	City        string
	Country     string
	Alias       []string
	Regions     []string
	Coordinates []float64
	Province    string
	Timezone    string
	Unlocs      []string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Port) Copy() *Port {
	if p == nil {
		return nil
	}
	return &Port{
		ID:          p.ID,
		Name:        p.Name,
		Code:        p.Code,
		City:        p.City,
		Country:     p.Country,
		Alias:       append([]string(nil), p.Alias...),
		Regions:     append([]string(nil), p.Regions...),
		Coordinates: append([]float64(nil), p.Coordinates...),
		Province:    p.Province,
		Timezone:    p.Timezone,
		Unlocs:      append([]string(nil), p.Unlocs...),
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func portStoreToDomain(p *Port) (*domain.Port, error) {
	if p == nil {
		return nil, errors.New("store port is nil")
	}
	return domain.NewPort(
		p.ID,
		p.Name,
		p.Code,
		p.City,
		p.Country,
		append([]string(nil), p.Alias...),
		append([]string(nil), p.Regions...),
		append([]float64(nil), p.Coordinates...),
		p.Province,
		p.Timezone,
		append([]string(nil), p.Unlocs...),
	)
}

func portDomainToStore(p *domain.Port) *Port {
	return &Port{
		ID:          p.ID(),
		Name:        p.Name(),
		Code:        p.Code(),
		City:        p.City(),
		Country:     p.Country(),
		Alias:       append([]string(nil), p.Alias()...),
		Regions:     append([]string(nil), p.Regions()...),
		Coordinates: append([]float64(nil), p.Coordinates()...),
		Province:    p.Province(),
		Timezone:    p.Timezone(),
		Unlocs:      append([]string(nil), p.Unlocs()...),
	}
}
