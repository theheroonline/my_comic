package main

import (
	"go_comic/routers"
	"net/http"
	"time"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           ":11036",
		Handler:        router,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
