package routes

import (
	"net/http"

	controllers "github.com/nicolasrsaraiva/allc-api/src/controller"
)

func InitRoutes() {
	http.HandleFunc("/user/create", controllers.CreateUserController)
}
