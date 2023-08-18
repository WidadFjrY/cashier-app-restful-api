package routes

import (
	"github.com/julienschmidt/httprouter"
	controller "github.com/widadfjry/cashier-app/controllers/product"
)

func ProductRoutes(router *httprouter.Router, controller controller.ProductController) {
	router.POST("/api/v1/products/create", controller.Create)
	router.GET("/api/v1/products/id/:productId", controller.GetProductById)
	router.PUT("/api/v1/products/id/:productId", controller.UpdateProductById)
	router.DELETE("/api/v1/products/id/:productId", controller.DeleteProductById)
	router.GET("/api/v1/products/pages/:page", controller.GetProducts)
}
