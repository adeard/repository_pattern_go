package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookrequest BookRequest) (Book, error)
	Update(ID int, bookrequest BookRequest) (Book, error)
	Delete(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {

	books, err := s.repository.FindAll()

	return books, err
}

func (s *service) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)

	return book, err
}

func (s *service) Create(bookrequest BookRequest) (Book, error) {

	price, err := bookrequest.Price.Int64()

	book, err := s.repository.Create(Book{
		Title: bookrequest.Title,
		Price: int(price),
	})

	return book, err
}

func (s *service) Update(ID int, bookrequest BookRequest) (Book, error) {

	book, err := s.repository.FindByID(ID)

	price, err := bookrequest.Price.Int64()

	book.Price = int(price)
	book.Title = bookrequest.Title

	newbook, err := s.repository.Update(book)

	return newbook, err
}

func (s *service) Delete(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	deletedBook, err := s.repository.Delete(book)

	return deletedBook, err
}
