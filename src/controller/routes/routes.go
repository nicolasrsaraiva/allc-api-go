package routes

import (
	"net/http"

	"github.com/nicolasrsaraiva/allc-api/src/controller"
)

func InitRoutes() {
	http.HandleFunc("/user", controller.CreateUserController)
	http.HandleFunc("/users", controller.GetUsersController)
}
