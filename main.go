package main

import (
	"fmt"
	"fun-w-flags/database"
	"fun-w-flags/routes"
	"net/http"
	"text/template"
)

// encapsula todos os templates
var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	fmt.Println("inicio do programa...")
	database.ConectaComBancoDeDados()

	routes.HandleRequests()

}

func index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", nil)
}
