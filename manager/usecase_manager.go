package manager

import "enigmacamp.com/golang-with-mongodb/usecase"

type UseCaseManager interface {
	ProductRegistrationUseCase() usecase.ProductRegistrationUseCase
}

type useCaseManager struct {
	repoManager RepositoryManager
}

func (u *useCaseManager) ProductRegistrationUseCase() usecase.ProductRegistrationUseCase {
	return usecase.NewProductRegistrationUseCase(u.repoManager.ProductRepo())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{repoManager: repoManager}
}
