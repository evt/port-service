package transport_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/evt/port-service/internal/repository/inmem"
	"github.com/evt/port-service/internal/services"
	"github.com/evt/port-service/internal/transport"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type HttpTestSuite struct {
	suite.Suite
	portService transport.PortService
	httpServer  transport.HttpServer
}

func NewHttpTestSuite() *HttpTestSuite {
	suite := &HttpTestSuite{}

	// create port repository
	portStoreRepo := inmem.NewPortStore()

	// create port service
	suite.portService = services.NewPortService(portStoreRepo)

	// create http server with application injected
	suite.httpServer = transport.NewHttpServer(suite.portService)

	return suite
}

func TestHttpTestSuite(t *testing.T) {
	suite.Run(t, NewHttpTestSuite())
}

func (suite *HttpTestSuite) TestUploadPorts() {
	portsRequest, err := os.ReadFile("testfixtures/ports_request.json")
	require.NoError(suite.T(), err)

	requestPortsTotal := countJSONPorts(suite.T(), portsRequest)

	portsResponse, err := os.ReadFile("testfixtures/ports_response.json")
	require.NoError(suite.T(), err)

	// create POST /ports request
	req := httptest.NewRequest(http.MethodPost, "/ports", bytes.NewBuffer(portsRequest))

	w := httptest.NewRecorder()

	// run request
	suite.httpServer.UploadPorts(w, req)

	res := w.Result()
	defer res.Body.Close()

	// read response body
	data, err := io.ReadAll(res.Body)
	require.NoError(suite.T(), err)

	require.Equal(suite.T(), http.StatusOK, res.StatusCode)
	require.Equal(suite.T(), portsResponse, data)

	// count ports in storage
	storedPortsTotal, err := suite.portService.CountPorts(context.Background())
	require.NoError(suite.T(), err)

	// compare number of ports in request and storage
	require.Equal(suite.T(), requestPortsTotal, storedPortsTotal)
}

func (suite *HttpTestSuite) TestUploadPorts_badJSON() {
	// create POST /ports request
	req := httptest.NewRequest(http.MethodPost, "/ports", bytes.NewBuffer([]byte("blabla")))

	w := httptest.NewRecorder()

	// run request
	suite.httpServer.UploadPorts(w, req)

	res := w.Result()
	defer res.Body.Close()

	require.Equal(suite.T(), http.StatusBadRequest, res.StatusCode)
}

// countJSONPorts counts the number of ports in provided JSON.
func countJSONPorts(t *testing.T, portsJSON []byte) int {
	t.Helper()
	var ports map[string]struct{}
	err := json.Unmarshal(portsJSON, &ports)
	require.NoError(t, err)
	return len(ports)
}
