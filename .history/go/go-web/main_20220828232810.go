package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Auto struct {
	ID     int    `json:"id"`
	Auto   string `json:"auto"`
	Modelo string `json:"modelo"`
	Año    int    `json:"año"`
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

	http.ListenAndServe(":9990", nil)
}
