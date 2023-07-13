package main

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Get("/json", handlerIPJSON)
	r.Get("/", handlerIPPianText)
	http.ListenAndServe(":3000", r)
}

func handlerIPJSON(w http.ResponseWriter, r *http.Request) {
	type msg struct {
		ID string `json:"ip"`
	}
	ip := msg{
		ID: getIP(w, r),
	}

	RespondWithJSON(w, r, 200, ip)
}

func handlerIPPianText(w http.ResponseWriter, r *http.Request) {
	ip := getIP(w, r)

	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte(ip))
}

func getIP(w http.ResponseWriter, r *http.Request) string {
	ip := r.RemoteAddr
	ip = strings.Split(ip, ":")[0]
	return ip

}
