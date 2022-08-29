package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type Auto struct {
	ID     int    `json:"id"`
	Auto   string `json:"auto"`
	Modelo string `json:"modelo"`
	Año    int    `json:"año"`
}

func main() {

	cntArch, err := ioutil.ReadFile("server-ubuntu/go/go-web/usuarios.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cntArch)
	/*
		c := Auto{}

		jsonSalida := json.Unmarshal(cntArch, &c)
	*/
	/*
		//jsonSalida := json.Unmarshal(cntArch, &c)

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		})

		http.ListenAndServe(":9990", nil)
	*/
}
