package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/souravToptal/ecom/internal/product"
)

type ProductHandler struct {
	service *product.ProductService
}

func NewProductHandler(service *product.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (handler *ProductHandler) CreateProduct(c echo.Context) (err error) {
	request := new(product.CreateProductRequest)
	if err = c.Bind(request); err != nil {
		return
	}

	err = handler.service.CreateProduct(request)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}

	return c.JSON(http.StatusCreated, nil)
}