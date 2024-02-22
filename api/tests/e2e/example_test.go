package e2e

import (
	"github.com/ellioht/go-rest-api/internal/server/mock_server"
	"testing"
)

func Test_Example(t *testing.T) {
	mockServer := mock_server.NewMockServer(nil)
	services := mockServer.SetupServices("issues")
	services.Issues.CreateIssue()
}
