package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/widadfjry/cashier-app/app"
	controller2 "github.com/widadfjry/cashier-app/controllers/category"
	controller3 "github.com/widadfjry/cashier-app/controllers/product"
	controller "github.com/widadfjry/cashier-app/controllers/user"
	"github.com/widadfjry/cashier-app/exception"
	"github.com/widadfjry/cashier-app/helper"
	"github.com/widadfjry/cashier-app/middleware"
	repositories2 "github.com/widadfjry/cashier-app/repositories/categories"
	repositories3 "github.com/widadfjry/cashier-app/repositories/products"
	repositories "github.com/widadfjry/cashier-app/repositories/user"
	"github.com/widadfjry/cashier-app/routes"
	services2 "github.com/widadfjry/cashier-app/services/category"
	services3 "github.com/widadfjry/cashier-app/services/product"
	services "github.com/widadfjry/cashier-app/services/user"
	"net/http"
	"time"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	serverRunning := color.New(color.FgBlack).Add(color.BgGreen).SprintfFunc()
	notifyRunning := []string{"Connect to database..", "Connected", "Starting server on port 3000..", "Server Listen and Serve"}
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	categoryRepository := repositories2.NewCategoryRepository()
	categoryService := services2.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller2.NewCategoryController(categoryService)

	productRepository := repositories3.NewProductRepository()
	productService := services3.NewProductService(productRepository, categoryRepository, db, validate)
	productController := controller3.NewProductController(productService)

	router := httprouter.New()

	routes.UserRouter(router, userController)
	routes.CategoryRouter(router, categoryController)
	routes.ProductRoutes(router, productController)

	router.PanicHandler = exception.ErrorHandle

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	for i := range notifyRunning {
		select {
		case <-ticker.C:
			fmt.Printf("%s ", serverRunning("OK"))
			fmt.Println(notifyRunning[i])
		}
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
