package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"twitter/server"
)

type tweet struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

func main() {
	s := server.Server{
		Repository: &server.TweetMemoryRepository{},
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/tweets", s.ListTweets)
	r.Post("/tweets", s.AddTweet)

	log.Fatalln(http.ListenAndServe(":8080", r))
}
