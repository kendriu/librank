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
	var titles []string
	for _, book := range s.repo.all(){
		titles = append(titles, book.Title)
	}
	return titles
}

func (s *ServiceImpl) Add(book Book) {
	s.repo.save(book)
}

func (s *ServiceImpl) Prune() {
	s.repo.prune()
}
