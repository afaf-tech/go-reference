package book

type Service interface {
	FindAll() ([]Book, error)
	FindOne(ID int) (Book, error)
	Create(book BookRequest) (Book, error)
	Update(ID int, book BookRequest) (Book, error)
	Delete(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repository: repo}
}

func (s *service) FindAll() ([]Book, error) {
	return s.repository.FindAll()
}

func (s *service) FindOne(ID int) (Book, error) {
	return s.repository.FindOne(ID)
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	bookPrice, _ := bookRequest.Price.Int64()
	discount, _ := bookRequest.Discount.Int64()
	rating, _ := bookRequest.Rating.Int64()
	book := Book{
		Title:       bookRequest.Title,
		Description: bookRequest.Description,
		Price:       int(bookPrice),
		Rating:      int(rating),
		Discount:    int(discount),
	}
	return s.repository.Create(book)
}

func (s *service) Update(ID int, bookRequest BookRequest) (Book, error) {
	book, err := s.repository.FindOne(ID)
	if err != nil {

	}
	bookPrice, _ := bookRequest.Price.Int64()
	discount, _ := bookRequest.Discount.Int64()
	rating, _ := bookRequest.Rating.Int64()

	book.Title = bookRequest.Title
	book.Description = bookRequest.Description
	book.Price = int(bookPrice)
	book.Rating = int(rating)
	book.Discount = int(discount)
	return s.repository.Update(ID, book)
}
func (s *service) Delete(ID int) (Book, error) {
	book, err := s.repository.FindOne(ID)
	if err != nil {

	}
	return s.repository.Delete(book)
}
