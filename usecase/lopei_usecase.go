package usecase

import (
	"livecode-wmb-rest-api/repository"
	"livecode-wmb-rest-api/service"
)

type LopeiUseCase interface {
	GetBalance(lopeId int32) (service.ResultMessage, error)
	DoPayment(lopeId int32, amount float32) (bool, error)
}

type lopeiUseCase struct {
	repo repository.CustRepo
}

func (c *lopeiUseCase) GetBalance(lopeId int32) (service.ResultMessage, error) {
	return c.repo.CheckBalance(lopeId)
}

func (c *lopeiUseCase) DoPayment(lopeId int32, amount float32) (bool, error) {
	return c.repo.DoPayment(lopeId, amount)
}

func NewLopeiUseCase(repo repository.CustRepo) LopeiUseCase {
	return &lopeiUseCase{repo: repo}
}
