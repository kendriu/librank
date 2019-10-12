package audioteka

type Service interface {
	GetTitles() []string
	Add(book Book)
	Prune()
}

type ServiceImpl struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &ServiceImpl{repo: repo}
}

func (s *ServiceImpl) GetTitles() []string {
	panic("implement me")
}

func (s *ServiceImpl) Add(book Book) {
	s.repo.save(book)
}

func (s *ServiceImpl) Prune() {
	s.repo.prune()
}
