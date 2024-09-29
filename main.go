package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World!"))
    })

    r.Post("/", func(w http.ResponseWriter, r *http.Request) {
        bodyData, err := io.ReadAll(r.Body)
      
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte("Body data is required"))
            return
        }

        w.WriteHeader(http.StatusOK)

        w.Write([]byte(fmt.Sprintf("Received: %s", bodyData)))
    })
    
    http.ListenAndServe(":3000", r)
}
