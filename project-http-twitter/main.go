package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
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

	go spamTweets()

	log.Fatalln(http.ListenAndServe(":8080", r))

}

func spamTweets() error {
	for {
		addTweetPayload := server.Tweet{
			Message:  "ass",
			Location: "ass",
		}
		marshaledPayload, err := json.Marshal(addTweetPayload)

		url := fmt.Sprintf("http://localhost:%v/tweets", os.Getenv("PORT"))

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(marshaledPayload))

		defer resp.Body.Close()

		if err != nil {
			return err
		}

	}
	return nil
}
