package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/evt/port-service/internal/domain"
)

func portHttpToDomain(p *Port) (*domain.Port, error) {
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

// readPorts reads ports from provided reader and sends them to portChan.
func readPorts(ctx context.Context, r io.Reader, portChan chan Port) error {
	decoder := json.NewDecoder(r)

	// Read opening delimiter
	t, err := decoder.Token()
	if err != nil {
		return fmt.Errorf("failed to read opening delimiter: %w", err)
	}

	// Make sure opening delimiter is `{`
	if t != json.Delim('{') {
		return fmt.Errorf("expected {, got %v", t)
	}

	for decoder.More() {
		// Check if context is cancelled.
		if ctx.Err() != nil {
			return ctx.Err()
		}
		// Read the port ID.
		t, err := decoder.Token()
		if err != nil {
			return fmt.Errorf("failed to read port ID: %w", err)
		}
		// Make sure port ID is a string.
		portID, ok := t.(string)
		if !ok {
			return fmt.Errorf("expected string, got %v", t)
		}

		// Read the port and send it to the channel.
		var port Port
		if err := decoder.Decode(&port); err != nil {
			return fmt.Errorf("failed to decode port: %w", err)
		}

		port.ID = portID
		portChan <- port
	}

	return nil
}
