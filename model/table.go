package model

import "gorm.io/gorm"

type Table struct {
	Id              	int `gorm:"primaryKey" json:"id"`
	TableDescription	string `json:"description" binding:"required"`
	IsAvailable     	bool `gorm:"default:true" json:"is_available"`
	BillID 				[]Bill
	gorm.Model
}

func (Table) TableName() string {
	return "m_table" //Ganti nama tabel
}