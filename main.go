package main

import (
	"net/http"

	"github.com/nicolasrsaraiva/allc-api/src/controller/routes"
)

func main() {

	routes.InitRoutes()
	http.ListenAndServe(":8080", nil)
}
