package usecase

import (
	"livecode-wmb-rest-api/model"
	"livecode-wmb-rest-api/repository"
	"livecode-wmb-rest-api/utils"
	"log"
)

type DailySalesUC interface {
	DailySales() ([]model.DtoDailySales, error)
}

type dailySalesUC struct {
	repo repository.BillRepo
}

func (u *dailySalesUC) DailySales() ([]model.DtoDailySales, error){
	sales, err := u.repo.DailySales()
	if utils.IsError(err) {
		log.Println("Failed at DailySales")
		return sales, err
	}
	return sales, nil
}

func NewDailySalesUseCase(repo repository.BillRepo) DailySalesUC {
	return &dailySalesUC{
		repo: repo,
	}
}