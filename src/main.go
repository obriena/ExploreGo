package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	fmt.Println("Starting Test Service")
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Returning welcome message")
		resp, err := http.Get("http://worldtimeapi.org/api/timezone/America/Chicago")
		defer resp.Body.Close()
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		w.Write([]byte(body))
	})

	r.Route("/v1", func(r chi.Router) {
		// Unauthenticated routes.
		route(r, "GET", "/health", http.HandlerFunc(Health))
		//		route(r, "GET", "/employer/{id}", http.HandlerFunc(handler.EmployerById(arcaClient)))
	})

	http.ListenAndServe(":8889", r)
}

func route(r chi.Router, method string, pattern string, handler http.Handler) {
	wrappedHandler := wrapHandler(pattern, handler)
	r.Method(method, pattern, wrappedHandler)
}

func wrapHandler(pattern string, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	})
}

func Health(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)
	m["message"] = "Everything seems okay :-)!"
	RespondJSON(w, r, m)
}

func RespondJSON(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	b, _ := json.Marshal(data)

	w.Write(b)
}
