package manager

import (
	"livecode-wmb-rest-api/repository"
	"livecode-wmb-rest-api/usecase"
)

type UseCaseManager interface {
	CustomerRegistrationUseCase() usecase.CustomerRegistrationUC
	CreateBillUseCase() usecase.CreateBillUC
	PrintBillUseCase() usecase.PrintBillUC
	DailySalesUseCase() usecase.DailySalesUC
	
	DiscCRUDUseCase() usecase.DiscCRUDUC
	MPriceCRUDUseCase() usecase.MPriceCRUDUC
	MenuCRUDUseCase() usecase.MenuCRUDUC
	TableCRUDUseCase() usecase.TableCRUDUC
	TransTypeCRUDUseCase() usecase.TransTypeCRUDUC

	LopeiUseCase() usecase.LopeiUseCase
	LoginUseCase() usecase.LoginUC
}

type useCaseManager struct {
	billRepo 	repository.BillRepo
	custRepo	repository.CustRepo
	menuRepo	repository.MenuRepo	
	loginRepo 	repository.LoginRepo
}

func (u *useCaseManager) CustomerRegistrationUseCase() usecase.CustomerRegistrationUC {
	return usecase.NewCustomerRegistrationUseCase(u.custRepo)
}

func (u *useCaseManager) CreateBillUseCase() usecase.CreateBillUC {
	return usecase.NewCreateBillUseCase(u.billRepo)
}

func (u *useCaseManager) PrintBillUseCase() usecase.PrintBillUC {
	return usecase.NewPrintBillUseCase(u.billRepo)
}

func (u *useCaseManager) DailySalesUseCase() usecase.DailySalesUC {
	return usecase.NewDailySalesUseCase(u.billRepo)
}

func (u *useCaseManager) DiscCRUDUseCase() usecase.DiscCRUDUC {
	return usecase.NewDiscCRUDUseCase(u.custRepo)
}

func (u *useCaseManager) MPriceCRUDUseCase() usecase.MPriceCRUDUC {
	return usecase.NewMPriceCRUDUseCase(u.menuRepo)
}

func (u *useCaseManager) MenuCRUDUseCase() usecase.MenuCRUDUC {
	return usecase.NewMenuCRUDUseCase(u.menuRepo)
}

func (u *useCaseManager) TableCRUDUseCase() usecase.TableCRUDUC {
	return usecase.NewTableCRUDUseCase(u.billRepo)
}

func (u *useCaseManager) TransTypeCRUDUseCase() usecase.TransTypeCRUDUC {
	return usecase.NewTransTypeCRUDUseCase(u.billRepo)
}

func (u *useCaseManager) LopeiUseCase() usecase.LopeiUseCase {
	return usecase.NewLopeiUseCase(u.custRepo)
}

func (u *useCaseManager) LoginUseCase() usecase.LoginUC {
	return usecase.NewLoginUC(u.loginRepo)
}

func NewUseCaseManager(billRepo repository.BillRepo, custRepo repository.CustRepo, menuRepo repository.MenuRepo, loginRepo repository.LoginRepo) UseCaseManager {
	return &useCaseManager{
		billRepo: billRepo,
		custRepo: custRepo,
		menuRepo: menuRepo,
		loginRepo: loginRepo,
	}
}