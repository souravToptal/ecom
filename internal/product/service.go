package product

import (
	"github.com/jinzhu/gorm"
)

type ProductService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{
		db: db,
	}
}

func (service *ProductService) CreateProduct(request *CreateProductRequest) (err error) {
	service.db.Create(&Product{Name: request.Name, Price: request.Price})
	return nil
}