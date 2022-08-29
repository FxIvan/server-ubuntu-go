package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Auto struct {
	ID     int    `json:"id"`
	Auto   string `json:"auto"`
	Modelo string `json:"modelo"`
	Año    int    `json:"año"`
}

func main() {

	cntArch, err := ioutil.ReadFile("usuarios.json")

	if err != nil {
		log.Fatal(err)
	}

	c := Auto{}

	jsonSalida := json.Unmarshal(cntArch, &c)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, jsonSalida.Error())
	})

	http.ListenAndServe(":9990", nil)
}
