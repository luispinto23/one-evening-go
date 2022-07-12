package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"sync"
	"twitter/server"
)

func main() {
	s := server.Server{
		Repository: &server.TweetMemoryRepository{
			Lock: &sync.RWMutex{},
		},
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/tweets", s.ListTweets)
	r.Post("/tweets", s.AddTweet)

	log.Fatalln(http.ListenAndServe(":8080", r))
}
