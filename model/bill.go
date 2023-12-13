package model

import (
	"time"

	"gorm.io/gorm"
)

type Bill struct {
	Id          	int `gorm:"primaryKey" json:"id"`
	TransDate   	time.Time
	CustomerID		uint `json:"customer_id" binding:"required"`
	TableID 		uint `json:"table_id"`
	TransTypeID		string `json:"trans_id" binding:"required"`
	BillDetailID	[]BillDetail `json:"bill_detail" binding:"required"`
	gorm.Model
}

type BillDetail struct {
	Id          int `gorm:"primaryKey" json:"id"`
	BillID      int `json:"bill_id"`
	MenuPriceID int `json:"mprice_id" binding:"required"`
	Quantity    int `json:"qty" binding:"required"`
	gorm.Model
}

type TransType struct {
	Id          string `gorm:"primaryKey" json:"id" binding:"required"`
	Description string `gorm:"not null" json:"description" binding:"required"`
	BillID 				[]Bill
	gorm.Model
}

func (Bill) TableName() string {
	return "t_bill" //Ganti nama tabel
}