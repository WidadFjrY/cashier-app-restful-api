package routes

import (
	"github.com/julienschmidt/httprouter"
	controller "github.com/widadfjry/cashier-app/controllers/user"
)

func UserRouter(router *httprouter.Router, userController controller.UserController) {
	// Public API
	router.POST("/api/v1/users/register", userController.Create)
	router.POST("/api/v1/users/login", userController.Login)
	router.POST("/api/v1/users/logout", userController.Logout)
	router.POST("/api/v1/users/verification", userController.Verification)
	router.POST("/api/v1/users/new-password", userController.NewPassword)

	// API
	router.PUT("/api/v1/users/current/change-password", userController.UpdatePassword)
	router.PUT("/api/v1/users/current/update-data", userController.UpdateByUsername)
	router.POST("/api/v1/users/add", userController.Add)

	// Rute with parameter wildcard
	router.GET("/api/v1/users/pages/:page", userController.GetAll)
	router.GET("/api/v1/users/id/:userId", userController.GetById)
	router.PUT("/api/v1/users/id/:userId", userController.UpdateById)
	router.DELETE("/api/v1/users/id/:userId", userController.DeleteById)
}
