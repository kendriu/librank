package lubimy_czytac

type Service interface {
	Update(*Book)
	Prune()
}

type ServiceImpl struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &ServiceImpl{repo: repo}
}

func (s *ServiceImpl) Update(book *Book) {
	repoBook :=s.repo.get(book.NeedleTitle)
	repoBook.updateItems(book.Items)
	s.repo.save(repoBook)
}

func (s *ServiceImpl) Prune() {
	s.repo.prune()
}
