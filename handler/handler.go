package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type RequestHandler interface {
	CanHandle(request *http.Request) bool
	Handle(w http.ResponseWriter, r *http.Request) error
}

type Router struct {
	handlers []RequestHandler
}

func (r *Router) AddHandler(handler RequestHandler) {
	r.handlers = append(r.handlers, handler)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, handler := range r.handlers {
		if handler.CanHandle(req) {
			if err := handler.Handle(w, req); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}
	http.NotFound(w, req)
}

func NewRouter() *Router {
	return &Router{}
}

type JSONHandler struct {
}

func (h *JSONHandler) CanHandle(request *http.Request) bool {
	return request.Header.Get("Content-Type") == "application/json"
}

// Handle обрабатывает JSON запрос
func (h *JSONHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	var data DataSet
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	// Обработка данных
	log.Printf("Received data: %+v", data)

	// Отправляем ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

// DataSet - структура для парсинга JSON
type DataSet struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}
