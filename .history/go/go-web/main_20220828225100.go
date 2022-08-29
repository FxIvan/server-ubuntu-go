package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Auto struct {
	ID     int
	Marca  string
	Modelo string
	Año    int
	Color  string
}

func main() {

	auto := Auto{
		ID:     1,
		Marca:  "Volkswagen",
		Modelo: "Siroco",
		Año:    2019,
		Color:  "Rojo",
	}

	crear_json, _ := json.Marshal(auto)

	convertir_cadena := string(crear_json)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, convertir_cadena)
	})

	http.ListenAndServer(":9990", nil)
}
