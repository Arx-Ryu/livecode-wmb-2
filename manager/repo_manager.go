package manager

import (
	"livecode-wmb-rest-api/repository"
)

type RepositoryManager interface {
	BillRepo() 		repository.BillRepo
	MenuRepo() 		repository.MenuRepo
	CustomerRepo() 	repository.CustRepo
	LoginRepo() 	repository.LoginRepo
}

type repositoryManager struct {
	infra	Infra
}

func (r *repositoryManager) BillRepo() repository.BillRepo {
	return repository.NewBillRepository(r.infra.SqlDb())
}

func (r *repositoryManager) MenuRepo() repository.MenuRepo {
	return repository.NewMenuRepository(r.infra.SqlDb())
}

func (r *repositoryManager) CustomerRepo() repository.CustRepo {
	return repository.NewCustomerRepository(r.infra.SqlDb(), r.infra.LopeiClientConn())
}

func (r *repositoryManager) LoginRepo() repository.LoginRepo {
	return repository.NewLoginRepository(r.infra.Token())
}

func NewRepoManager(infra Infra) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}