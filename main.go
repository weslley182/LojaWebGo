package main

import (
	"net/http"

	"LojaWeb/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
