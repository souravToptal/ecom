package product

import (
	"strconv"

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
	service.db.First(&product, id)
	return product, nil

}

func (service *ProductService) UpdateProduct(request map[string]interface{}) (Product, error) {
	var product Product
	id, err := strconv.Atoi(request["id"].(string))
	if err != nil {
		return product, err
	}
	service.db.First(&product, id)
	product.Name = request["name"].(string)
	product.Price = uint(request["price"].(float64))
	service.db.Model(&product).Where("id = ?", id).Update(&product)
	return product, nil
}

func (service *ProductService) DeleteProduct(id int) error {
	var product Product
	service.db.First(&product, id)
	service.db.Delete(&product)
	return nil
}
