package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shudipta/coding-test/api"
	"github.com/shudipta/coding-test/model"
	"github.com/shudipta/coding-test/repository"
)

// @title Sample CRUD api for test db
// @version 1.0
// @description Sample CRUD api for test db
// @termsOfService

// @contact.name Me
// @contact.url http://me.com/terms.html
// @contact.email me@me.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}

	db.LogMode(true)
	repository.DB = db

	db.AutoMigrate(
		&model.Category{},
		&model.Customer{},
		&model.OrderDetail{},
		&model.Order{},
		&model.Product{},
	)

	runServer()
}

func runServer() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", api.Hello)

	e.PATCH("/set-rating", api.SetRating)
	e.GET("/top-categories", api.FetchTopCategories)

	// Start server
	e.Logger.Fatal(e.Start(":80"))
}
