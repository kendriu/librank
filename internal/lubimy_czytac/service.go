package lubimy_czytac

type Service interface {
	Add(Book)
	Prune()
}

type ServiceImpl struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &ServiceImpl{repo: repo}
}

func (s *ServiceImpl) Add(Book) {
	panic("implement me")
}

func (s *ServiceImpl) Prune() {
	s.repo.prune()
}
