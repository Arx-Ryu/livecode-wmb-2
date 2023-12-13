package usecase

import (
	"livecode-wmb-rest-api/model"
	"livecode-wmb-rest-api/repository"
	"livecode-wmb-rest-api/utils"
	"log"
)

type CreateBillUC interface {
	CreateBill(newBill *model.Bill) error
}

type createBillUC struct {
	repo repository.BillRepo
}

func (u *createBillUC) CreateBill(newBill *model.Bill) error {	
	table, err := u.repo.FindByIdTable(int(newBill.TableID))
	if table.TableDescription == "" {
		log.Println("Failed to create bill, table does not exist!")
		return err
	} else if !table.IsAvailable {
		log.Println("Failed to create bill, table is not available!")
		return err
	}else if utils.IsError(err) {
		log.Println("Failed to create bill, error at findbyid table")
		return err
	}	
	err = u.repo.Create(newBill)
	if utils.IsError(err) {
		log.Println("Failed at Create Bill")
		return err
	}
	err = u.repo.UpdateTableAvailability(&table, !table.IsAvailable)
	if utils.IsError(err) {
		log.Println("Failed at Update Table Availability")
		return err
	}
	return nil
}

func NewCreateBillUseCase(repo repository.BillRepo) CreateBillUC {
	return &createBillUC{
		repo: repo,
	}
}