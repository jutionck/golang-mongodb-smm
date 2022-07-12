package usecase

import (
	"enigmacamp.com/golang-with-mongodb/model"
	"enigmacamp.com/golang-with-mongodb/repository"
)

type ProductRegistrationUseCase interface {
	Register(newProduct *model.Product) error
}

type productRegistrationUseCase struct {
	repo repository.ProductRepository
}

func (p *productRegistrationUseCase) Register(newProduct *model.Product) error {
	return p.repo.Add(newProduct)
}

func NewProductRegistrationUseCase(repo repository.ProductRepository) ProductRegistrationUseCase {
	return &productRegistrationUseCase{repo: repo}
}
