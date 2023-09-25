package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func HoraCertaHandler(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05")
	log.Println("chamou a API hora Certa")
	fmt.Fprintf(w, "<h1>Hora certa: %s<h1>", s)
}

func mainn() {
	http.HandleFunc("/horaCerta", HoraCertaHandler)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
