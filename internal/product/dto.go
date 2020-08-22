package product

type CreateProductRequest struct {
	Name    string `json:"name"`
	Price   uint   `json:"price"`
}