package main

import (
	"encoding/json"
	"fmt"
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

	cntArch, err := ioutil.ReadFile("C:/Users/almen/OneDrive/Escritorio/server-ubuntu/go/go-web/usuarios.json")

	if err != nil {
		log.Fatal(err)
	}

	c := Auto{}

	jsonSalida := json.Unmarshal(cntArch, &c)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Pritln(w, jsonSalida)
	})

	http.ListenAndServe(":9990", nil)
}