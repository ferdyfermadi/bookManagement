package bookmanagement

import (
	"bookManagement/handler"
	"bookManagement/router"
	"bookManagement/storage"
	"net/http"
)

func main() {
	bookStore := storage.GetBookStoreInstance()

	bookHandler := handler.NewBookHandler(bookStore)

	r := router.NewRouter(bookHandler)

	http.ListenAndServe(":8080", r)
}
