package handler

import (
	"bookManagement/model"
	"bookManagement/storage"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type BookHandler struct {
	store *storage.BookStore
}

func NewBookHandler(store *storage.BookStore) *BookHandler {
	return &BookHandler{store: store}
}

func (h *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := h.store.GetAll()
	writeJSON(w, http.StatusOK, books)
}

func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	book, found := h.store.GetByID(strconv.Itoa(id))
	if !found {
		writeError(w, http.StatusNotFound, "Book not found")
		return
	}

	writeJSON(w, http.StatusOK, book)
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	created := h.store.Create(book)
	writeJSON(w, http.StatusCreated, created)
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var book model.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	updated, ok := h.store.Update(strconv.Itoa(id), book)
	if !ok {
		writeError(w, http.StatusNotFound, "Book not found")
		return
	}

	writeJSON(w, http.StatusOK, updated)
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	ok := h.store.Delete(strconv.Itoa(id))
	if !ok {
		writeError(w, http.StatusNotFound, "Book not found")
		return
	}

	writeJSON(w, http.StatusNoContent, nil)
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}
