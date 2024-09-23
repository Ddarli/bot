package product

import (
	"errors"
	"strconv"
)

type Service struct{}

type ServiceInt interface {
	List()
	Create() *Product
	Delete(id int)
	Update(id int, title string)
	Read(id int)
}

func NewService() *Service {
	return &Service{}
}

func (*Service) List() []Product {
	return allProducts
}

func (*Service) Create() *Product {
	title := "Product - " + strconv.Itoa(len(allProducts)+1)
	newProduct := Product{Title: title}
	allProducts = append(allProducts, newProduct)
	return &newProduct
}

func (*Service) Get(id int) (*Product, error) {
	if id-1 < 0 || id-1 >= len(allProducts) {
		return nil, errors.New("product not found")
	}
	return &allProducts[id-1], nil
}
func (*Service) Delete(id int) (bool, error) {
	if id-1 < 0 || id-1 >= len(allProducts) {
		return false, errors.New("product not found")
	}
	removeIndex(id - 1)
	return true, nil
}
func (*Service) Update(id int, title string) (*Product, error) {
	if id-1 < 0 || id-1 >= len(allProducts) {
		return nil, errors.New("product not found")
	}
	allProducts[id-1].Title = title
	return &allProducts[id-1], nil
}
func removeIndex(index int) {
	allProducts = append(allProducts[:index], allProducts[index+1:]...)
}
