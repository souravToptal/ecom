package product

import (
	"github.com/jinzhu/gorm"
)

//ProductService ...
type ProductService struct {
	db *gorm.DB
}

//NewProductService ...
func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{
		db: db,
	}
}

//CreateProduct ...
func (service *ProductService) CreateProduct(request *CreateProductRequest) (Product, error) {
	product := Product{Name: request.Name, Price: request.Price}
	service.db.Create(&product)
	return product, nil
}

//GetProduct ...
func (service *ProductService) GetProduct(id int) (Product, error) {
	var product Product
	service.db.Select("name", "price").First(&product, id)
	return product, nil

}
