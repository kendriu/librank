package audioteka

type Service interface {
	GetUrls([]Book)
	Add(book Book)
	Prune()
}

type ServiceImpl struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &ServiceImpl{repo: repo}
}

func (s *ServiceImpl) GetUrls([]Book) {
	panic("implement me")
}

func (s *ServiceImpl) Add(book Book) {
	s.repo.save(book)
}

func (s *ServiceImpl) Prune() {
	s.repo.prune()
}
