package storage

import (
	"bookManagement/model"
	"sync"

	"github.com/google/uuid"
)

type BookStore struct {
	mu    sync.RWMutex
	books map[string]model.Book
}

var instance *BookStore
var once sync.Once

func GetBookStoreInstance() *BookStore {
	once.Do(func() {
		instance = &BookStore{
			books: make(map[string]model.Book),
		}
	})
	return instance
}

func (bs *BookStore) GetAll() []model.Book {
	bs.mu.RLock()
	defer bs.mu.RUnlock()
	result := make([]model.Book, 0, len(bs.books))
	for _, book := range bs.books {
		result = append(result, book)
	}
	return result
}

func (bs *BookStore) GetByID(id string) (model.Book, bool) {
	bs.mu.RLock()
	defer bs.mu.RUnlock()
	book, exists := bs.books[id]
	return book, exists
}

func (bs *BookStore) Create(book model.Book) model.Book {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	book.ID = uuid.New().String()
	bs.books[book.ID] = book
	return book
}

func (bs *BookStore) Update(id string, book model.Book) (model.Book, bool) {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	_, exists := bs.books[id]
	if !exists {
		return model.Book{}, false
	}
	book.ID = id
	bs.books[id] = book
	return book, true
}

func (bs *BookStore) Delete(id string) bool {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	_, exists := bs.books[id]
	if !exists {
		return false
	}
	delete(bs.books, id)
	return true
}
