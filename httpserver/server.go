package httpserver

import (
	"context"
	"log"
	"net/http"
)

func HandleFunc() {
	http.HandleFunc("/user", AuthMiddleware(LoggerMiddleware(CheckUser)))
	http.ListenAndServe(":8080", nil)
}

func CheckUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("id").(string)
		log.Println(user)
		next(w, r) // передает управление следующему middleware

	}
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, "id", "Получены данные из контекста")
		next(w, r.WithContext(ctx)) // передает управление следующему middleware с контекстом
	}
}

//Этой командой можно просмотреть данные в текущем эндпоните
// curl -v localhost:8080/user
