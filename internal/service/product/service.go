package product

import "errors"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (*Service) List() []Product {
	return allProducts
}

func (*Service) GetById(id int) (*Product, error) {
	if id < 0 || id >= len(allProducts) {
		return nil, errors.New("product not found")
	}
	return &allProducts[id], nil
}
