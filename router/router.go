package router

import (
	"bookManagement/handler"
	"bookManagement/middleware"

	"github.com/go-chi/chi/v5"
)

func NewRouter(h *handler.BookHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/books", func(r chi.Router) {
		r.Get("/", h.GetAllBooks)
		r.Get("/{id}", h.GetBookByID)
		r.Post("/", h.CreateBook)
		r.Put("/{id}", h.UpdateBook)
		r.Delete("/{id}", h.DeleteBook)
	})

	return r
}
