package issues

type MockService struct{}

func NewMockService() *MockService {
	return &MockService{}
}

func (s *MockService) CreateIssue() {

}
