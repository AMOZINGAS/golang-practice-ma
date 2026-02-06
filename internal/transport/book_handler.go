package transport

import (
	"encoding/json"
	"golang-practice-ma/internal/model"
	"golang-practice-ma/internal/service"
	"net/http"
	"strconv"
	"strings"
)

type Bookhandler struct {
	service *service.Service
}

func New(s *service.Service) *Bookhandler {
	return &Bookhandler{
		service: s,
	}
}

func (h *Bookhandler) HandleBooks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		libros, err := h.service.ObtenerTodosLosLibros()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(libros)

	case http.MethodPost:
		var libro model.Libro
		if err := json.NewDecoder(r.Body).Decode(&libro); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		created, err := h.service.CrearLibro(libro)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(created)
	default:
		http.Error(w, "Metodo no disponible", http.StatusMethodNotAllowed)
	}
}

func (h *Bookhandler) HandleBooksByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "No lo encontre", http.StatusBadRequest)
	}

	switch r.Method {

	case http.MethodGet:
		libro, err := h.service.ObtenerLibrosPorID(id)
		if err != nil {
			http.Error(w, "No lo encontramos", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(libro)

	case http.MethodPut:
		var libro model.Libro
		if err := json.NewDecoder(r.Body).Decode(&libro); err != nil {
			http.Error(w, "Input invalido", http.StatusBadRequest)
			return
		}

		updated, err := h.service.ActualizarLibro(id, libro)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updated)

	case http.MethodDelete:
		if err := h.service.EliminarLibro(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Metodo no disponible", http.StatusMethodNotAllowed)
	}
}
