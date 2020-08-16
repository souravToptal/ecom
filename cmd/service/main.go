package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    // "github.com/lib/pq"
)

func main () {
	// Echo instance
	e := echo.New()

	//DB Connection
	const addr = "postgresql://root@localhost:26257/ecom?sslmode=disable"
    db, err := gorm.Open("postgres", addr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

	fmt.Println(db)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
  
	// Routes
	e.GET("/", hello)
  
	// Start server
	e.Logger.Fatal(e.Start(":1324"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}