package product

//CreateProductRequest request body...
type CreateProductRequest struct {
	Name  string `json:"name"`
	Price uint   `json:"price"`
}
