package transport

import (
	"context"
	"net/http"

	"github.com/evt/port-service/internal/common/server"
	"github.com/evt/port-service/internal/domain"
)

// PortService is a port service
type PortService interface {
	GetPort(ctx context.Context, id string) (*domain.Port, error)
}

// HttpServer is a HTTP server for ports
type HttpServer struct {
	service PortService
}

// NewHttpServer creates a new HTTP server for ports
func NewHttpServer(service PortService) HttpServer {
	return HttpServer{
		service: service,
	}
}

// GetPort returns a port by ID
func (h HttpServer) GetPort(w http.ResponseWriter, r *http.Request) {
	port, err := h.service.GetPort(r.Context(), r.URL.Query().Get("id"))
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	response := Port{
		ID:          port.ID(),
		Name:        port.Name(),
		City:        port.City(),
		Country:     port.Country(),
		Alias:       port.Alias(),
		Regions:     port.Regions(),
		Coordinates: port.Coordinates(),
		Province:    port.Province(),
		Timezone:    port.Timezone(),
		Unlocs:      port.Unlocs(),
	}

	server.RespondOK(response, w, r)
}
