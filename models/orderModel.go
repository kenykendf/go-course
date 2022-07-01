package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CustomerName string `json:"customer_name" gorm:"type:varchar(100);column:customer_name"`
	OrderedAt    string `json:"ordered_at" gorm:"type:varchar(50);column:ordered_at"`
	Items        []Item `json:"items"`
}
