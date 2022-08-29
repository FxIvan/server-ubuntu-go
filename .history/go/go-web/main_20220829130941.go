package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Concesionaria struct {
	Autos []Auto `json:"concesionarioAutos"`
}

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

	c := Concesionaria{}

	_ = json.Unmarshal([]byte(cntArch), &c)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Content-Type", "application/json")
		//json.NewEncoder(w).Encode(cntArch)
		log.Panicln(c.Autos)
	})

	http.ListenAndServe(":9990", nil)
	//jsonSalida := json.Unmarshal(cntArch, &c)

	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			for i := 0; i < len(c.Autos); i++ {
				//fmt.Println("Product Id: ", c.Autos[i])
				io.WriteString(w, c.Autos[i])
			}
		})

		http.ListenAndServe(":9990", nil)
	*/
}
