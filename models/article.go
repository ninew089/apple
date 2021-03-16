package models

import (
	"github.com/jinzhu/gorm"
	pq "github.com/lib/pq"
)

type Article struct {
	gorm.Model
	Title   string `gorm:"unique;not null"`
	Category string `gorm:"not null"`
	Price string `gorm:"not null"`
	Body    string `gorm:"not null"`
	Image   string `gorm:"not null"`
	
}
type Customer struct {
	gorm.Model
	Name   string `gorm:"not null"`
	Tel string `gorm:"not null"`
	Address string `gorm:"not null"`
	Paid int `gorm:"not null"`
	ProductId  pq.Int64Array `gorm:"type:integer[]"`
}
