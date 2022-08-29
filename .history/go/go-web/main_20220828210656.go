package main

import "encoding/json"

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

	crear_json, _ = json.Marshal(auto)

}
