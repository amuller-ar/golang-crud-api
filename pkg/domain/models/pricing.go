package models

import "gorm.io/gorm"

type Pricing struct {
	gorm.Model
	SalePrice         float64
	AdministrativeFee float64
}
