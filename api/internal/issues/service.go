package issues

type IIssuesDB interface {
}

type Service struct {
	db IIssuesDB
}

func NewService(db IIssuesDB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) CreateIssue() {

}
