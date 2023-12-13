package usecase

import (
	"livecode-wmb-rest-api/model"
	"livecode-wmb-rest-api/repository"
	"livecode-wmb-rest-api/utils"
	"log"
	"strconv"
)

type PrintBillUC interface {
	PrintBill(id int) (model.DtoReceipt, error)
}

type printBillUC struct {
	repo repository.BillRepo
}

func (u *printBillUC) PrintBill(id int) (model.DtoReceipt, error) {
	var receipt model.DtoReceipt
	var print model.DtoBillNominal
	total := 0
	bill, err := u.repo.FindById(id)
	if utils.IsError(err) {
		log.Println("Failed at FindById Bill")
		return receipt, err
	}
	billId := strconv.Itoa(bill.Id)
	details, err := u.repo.FindSumBill(billId)
	if utils.IsError(err) {
		log.Println("Failed at FindAllBillDetailByBillId")
		return receipt, err
	}
	for _, items := range details {
		total = total + (items.Quantity * items.Price)
	}
	print = model.DtoBillNominal{
		Subtotal: total,
	}
	disc, err := u.repo.FindBillDisc(int(bill.CustomerID))
	if utils.IsError(err) {
		log.Println("Failed at FindBillDisc")
		return receipt, err
	}
	for _, items := range disc {
		total = total - (total * items.Pct / 100)
	}
	print = model.DtoBillNominal{
		Subtotal: print.Subtotal,
		Total: total,
	}
	table, err := u.repo.FindByIdTable(int(bill.TableID))
	if utils.IsError(err) {
		log.Println("Failed at Findbyid table")
		return receipt, err
	}
	err = u.repo.UpdateTableAvailability(&table, !table.IsAvailable)
	if utils.IsError(err) {
		log.Println("Failed at Update Table Availability")
		return receipt, err
	}
	receipt = model.DtoReceipt{
		Bill: bill,
		DtoBillNominal: print,
	}
	return receipt, nil
}

func NewPrintBillUseCase(repo repository.BillRepo) PrintBillUC {
	return &printBillUC{
		repo: repo,
	}
}