package product

import (
	"github.com/jinzhu/gorm"
)

//Product Schema...
type Product struct {
	gorm.Model
	Name  string
	Price uint
}
