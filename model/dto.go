package model

type DtoBillDetailMenuPrice struct {
	Quantity int
	Price    int
}

type DtoDailySales struct {
	Date		string
	Total		string
}

type DtoBillNominal struct {
	Subtotal 	int
	Total		int
}

type DtoReceipt struct {
	Bill
	DtoBillNominal
}

type DtoPaymentLopei struct {
	BillId 		int
	LopeiId		int
	Amount 		float32
}