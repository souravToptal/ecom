package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/souravToptal/ecom/cmd/service/handlers"
	"github.com/souravToptal/ecom/internal/product"
)

func main() {
	// Echo instance
	e := echo.New()

	//DB Connection
	const addr = "postgresql://root@localhost:26257/ecom?sslmode=disable"
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&product.Product{})

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	productService := product.NewProductService(db)
	productHandler := handlers.NewProductHandler(productService)

	// Routes
	e.GET("/health", health)
	e.POST("/products", productHandler.CreateProduct)
	e.GET("/products/:id", productHandler.GetProduct)
	e.PUT("/products", productHandler.UpdateProduct)
	e.DELETE("/products/:id", productHandler.DeleteProduct)

	// Start server
	e.Logger.Fatal(e.Start(":1324"))
}

// Handler
func health(c echo.Context) error {
	return c.String(http.StatusOK, "{\"Status\":\"OK\"}")
}
