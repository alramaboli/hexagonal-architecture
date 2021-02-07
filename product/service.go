package product

//Service ...
type Service interface {
	Find(code string) (*Product, error)
	Store(product *Product) error
	FindAll() ([]*Product, error)
	Delete(code string) error
}

//Repository ...
type ProductRepository interface {
	Find(code string) (*Product, error)
	Store(product *Product) error
	FindAll() ([]*Product, error)
	Delete(code string) error
}

type service struct {
	productrepo ProductRepository
}

//NewProductService ...
func NewProductService(productrepo ProductRepository) Service {

	return &service{productrepo: productrepo}
}

func (s *service) Find(code string) (*Product, error) {

	return s.productrepo.Find(code)

}

func (s *service) Store(product *Product) error {

	return s.productrepo.Store(product)

}

func (s *service) FindAll() ([]*Product, error) {
	return s.productrepo.FindAll()
}

func (s *service) Delete(code string) error {

	return s.productrepo.Delete(code)
}
