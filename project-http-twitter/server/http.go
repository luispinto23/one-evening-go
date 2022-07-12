package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

type tweet struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

type tweetResponse struct {
	ID int `json:"ID"`
}

type tweetsList struct {
	Tweets []tweet `json:"tweets"`
}

type TweetMemoryRepository struct {
	Tweets []tweet
	Lock   *sync.RWMutex
}

type Server struct {
	Repository *TweetMemoryRepository
}

func (r *TweetMemoryRepository) ListTweets() ([]tweet, error) {
	r.Lock.RLock()
	defer r.Lock.RUnlock()

	return r.Tweets, nil
}

//func (s Server) Tweets(w http.ResponseWriter, r *http.Request) {
//	if r.Method == http.MethodPost {
//		s.addTweet(w, r)
//	} else if r.Method == http.MethodGet {
//		s.listTweets(w, r)
//	}
//}

func (r *TweetMemoryRepository) AddTweet(m tweet) (int, error) {
	r.Lock.Lock()
	defer r.Lock.Unlock()

	r.Tweets = append(r.Tweets, m)
	return len(r.Tweets), nil
}

func (s Server) ListTweets(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	tweets, err := s.Repository.ListTweets()
	if err != nil {
		log.Println("Failed to add message:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tweetsList := tweetsList{
		Tweets: tweets,
	}

	response, err := json.Marshal(tweetsList)

	if err != nil {
		log.Println("Failed to marshal tweets:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer func() {
		duration := time.Since(start)
		fmt.Printf("%s %s %s\n", r.Method, r.URL, duration)
	}()

	w.Write(response)
}

func (s Server) AddTweet(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println("Failed to read body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	m := tweet{}

	if err := json.Unmarshal(body, &m); err != nil {
		log.Println("Failed to unmarshal payload:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := s.Repository.AddTweet(m)
	if err != nil {
		log.Println("Failed to add message:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Printf("Tweet: `%s` from %s with ID: %v \n", m.Message, m.Location, id)

	tweetID := tweetResponse{
		ID: id,
	}

	response, err := json.Marshal(tweetID)

	if err != nil {
		log.Println("Failed to marshal ID:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer func() {
		duration := time.Since(start)
		fmt.Printf("%s %s %s\n", r.Method, r.URL, duration)
	}()

	w.Write(response)
}
