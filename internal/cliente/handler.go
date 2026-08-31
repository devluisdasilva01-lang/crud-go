package cliente

import (
	"ecoding/json"
	"net/http"
	"strconv"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *Service 
}

func NetHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) addCliente(response http.ResponseWriter, request *http.Request) {
	
	var cliente Cliente

	err := json.NewDecoder(request.Body).Decode(&cliente)

	if err != nil {
		http.Error(
			response, "JSON invalido",
			http.StatusBadRequest,
		)

		return
	}

	err = h.service.CadastrarCliente(cliente)

	if err != nil {
		http.Error(
			response, err.Error(),
			http.StatusBadResquest,
		)

		return
	}

	response.WriteHeader(http.StatusCreated)
}

func(h *Handler) ListarTodosClientes(response http.ResponseWriter, request *http.Request) {

	clientes, err := h.service.ListarClientes()

	if err != nil {
		http.Error(
			response, "Error ao buscar cliente",
			http.StatusInternalServerError,
		)

		return
	}

	response.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(response).Enconde(clientes)
}

func(h *Handler) BuscarClientePorId (response http.ResponseWriter, request *http.Request) {

	idTexto := chi.UrlParam(request, "id")
	id, err : stroconv.Atoi(idText) 

	if err != nil {
		http.Error(
			response, "ID Inválido",
			http.StatusBadRequest,
		)
	}

	cliente, err := h.service.BuscarClientePorId(id)

	if err != nil {
		htt.Error(
			reponse, "Cliente não encontrado"
			http.StatusNotFound,
		)
	}

	json.NewEnconder(response).Encode(cliente)
}

func(h *Handler) EditarCliente(response http.ResponseWriter, request *http.Request) {
	
}