package routes

import (
	"github.com/julienschmidt/httprouter"
	controller "github.com/widadfjry/cashier-app/controllers/category"
)

func CategoryRouter(router *httprouter.Router, controller controller.CategoryController) {
	router.GET("/api/v1/categories/", controller.GetAllCategory)
	router.POST("/api/v1/categories/create", controller.Create)
	router.GET("/api/v1/categories/id/:categoryId", controller.GetCategoryById)
	router.PUT("/api/v1/categories/id/:categoryId", controller.UpdateById)
	router.DELETE("/api/v1/categories/id/:categoryId", controller.Delete)
}
