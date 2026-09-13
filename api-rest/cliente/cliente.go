package cliente

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)


type Cliente struct {
	ID int `json:"id"`
	Nome string `json:"nome"` 
	Telefone string `json:"telefone"`
	Email string `json:"email"`
}

type repositorio struct {
	mu	sync.Mutex
	clientes	map[int]Cliente
	proximoID int
}

var repo = &repositorio{
	clientes: make(map[int]Cliente),
	proximoID: 1,
}


func Criar(w http.ResponseWriter, r *http.Request) {
	var c Cliente
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if c.Nome == "" {
		http.Error(w, "o campo 'nome' é obrigatório", http.StatusBadRequest)
		return
	}

	repo.mu.Lock()
	c.ID = repo.proximoID
	repo.proximoID++
	repo.clientes[c.ID] = c
	repo.mu.Unlock()

	responderJSON(w, http.StatusCreated, c)
}

func Listar(w http.ResponseWriter, r *http.Request) {
	repo.mu.Lock()
	lista := make([]Cliente, 0, len(repo.clientes))
	for _, c := range repo.clientes {
		lista = append(lista, c)
	}
	repo.mu.Unlock()

	responderJSON(w, http.StatusOK, lista)
}

func BuscarPorID(w http.ResponseWriter, r *http.Request) {
	id, err := idDaURL(r)
	if err !=nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	repo.mu.Lock()
	c, ok := repo.clientes[id]
	repo.mu.Unlock()

	if !ok {
		http.Error(w, "cliente não encontrado", http.StatusNotFound)
		return
	}

	responderJSON(w, http.StatusOK, c)
}

func Atualizar(w http.ResponseWriter, r *http.Request) {
	id, err := idDaURL(r)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var dados Cliente
	if err := json.NewDecoder(r.Body).Decode(&dados); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	repo.mu.Lock()
	defer repo.mu.Unlock()

	if _, ok := repo.clientes[id]; !ok {
		http.Error(w, "cliente não encontrado", http.StatusNotFound)
		return
	}

	dados.ID = id
	repo.clientes[id] = dados
	
	responderJSON(w, http.StatusOK, dados)

}

func Remover(w http.ResponseWriter, r *http.Request) {
	id, err := idDaURL(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	repo.mu.Lock()
	defer repo.mu.Unlock()

	if _, ok := repo.clientes[id]; !ok {
		http.Error(w, "Cliente não encontrado", http.StatusNotFound)
		return
	}

	delete(repo.clientes, id)
	w.WriteHeader(http.StatusNoContent)
}

func idDaURL(r *http.Request) (int,error) {
	valor := mux.Vars(r)["id"]
	id, err := strconv.Atoi(valor)
	if err != nil {
		return 0, errors.New("id inválido, deve ser um número")
	}
	return id, nil
}

func responderJSON(w http.ResponseWriter, status int, corpo any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(corpo)
}