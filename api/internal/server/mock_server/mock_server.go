package mock_server

import (
	"github.com/ellioht/go-rest-api/internal/issues"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MockServer struct {
	Services *ServiceRegistry
	Deps     *dependencies
}

type dependencies struct {
	Issues *issues.Service
}

type ServiceList struct {
	Issues *issues.Service
}

func NewMockServer(pool *pgxpool.Pool) *MockServer {
	var deps dependencies

	registry := NewServiceRegistry()

	server := &MockServer{
		Services: registry,
	}

	// mocks

	// issues service
	issuesDb := issues.NewDb(pool)
	issuesService := issues.NewService(issuesDb)
	deps.Issues = issuesService
	if err := server.Services.Register(deps.Issues); err != nil {
		panic(err)
	}

	server.Deps = &deps
	return server
}

func (s *MockServer) SetupServices(services ...string) *ServiceList {
	serviceList := &ServiceList{}
	for _, service := range services {
		switch service {
		case "issues":
			err := s.Services.Fetch(&s.Deps.Issues)
			if err != nil {
				panic(err)
			}
			serviceList.Issues = s.Deps.Issues
		}
	}

	return serviceList
}
