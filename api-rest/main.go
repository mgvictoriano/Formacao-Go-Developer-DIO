package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"Formacao-Go-Developer-DIO/api-rest/cliente"
)

func main() {
r := mux.NewRouter()

r.HandleFunc("/clientes", cliente.Criar).Methods(http.MethodPost)
r.HandleFunc("/clientes", cliente.Listar).Methods(http.MethodGet)
r.HandleFunc("/clientes/{id}", cliente.BuscarPorID).Methods(http.MethodPut)
r.HandleFunc("/clientes/{id}", cliente.Remover).Methods(http.MethodDelete)

log.Println("servidor escutando em http://localhost:8080")
log.Fatal(http.ListenAndServe(":8080", r))


}