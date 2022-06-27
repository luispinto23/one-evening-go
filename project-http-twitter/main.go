package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

//type userPayload struct {
//	Name string `json:"name"`
//	Age  int    `json:"age"`
//}

type tweet struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

type tweetResponse struct {
	ID int `json:"ID"`
}

type server struct {
	repository *TweetMemoryRepository
}

type TweetMemoryRepository struct {
	tweets []tweet
}

type tweetsList struct {
	Tweets []tweet `json:"tweets"`
}

func (r *TweetMemoryRepository) AddTweet(m tweet) (int, error) {
	r.tweets = append(r.tweets, m)
	fmt.Println("TWEETS : %v", r.tweets)
	return len(r.tweets), nil
}

func (r *TweetMemoryRepository) Tweets() ([]tweet, error) {
	return r.tweets, nil
}

func (s server) tweets(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		s.addTweet(w, r)
	} else if r.Method == http.MethodGet {
		s.listTweets(w, r)
	}
}

func main() {
	s := server{
		repository: &TweetMemoryRepository{},
	}

	http.HandleFunc("/tweets", s.tweets)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func (s server) listTweets(w http.ResponseWriter, r *http.Request) {

	fmt.Println("IN THE GET METHOD!")
	tweets, err := s.repository.Tweets()
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

	w.Write(response)
}

func (s server) addTweet(w http.ResponseWriter, r *http.Request) {
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

	id, err := s.repository.AddTweet(m)
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

	w.Write(response)
}

//func myHandler(w http.ResponseWriter, r *http.Request) {
//	body, err := io.ReadAll(r.Body)
//	if err != nil {
//		log.Println("Failed to read body:", err)
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//	defer r.Body.Close()
//
//	u := userPayload{}
//	// Notice how we pass a pointer to the payload to json.Unmarshal.
//	// This is because the function modifies the struct
//	if err := json.Unmarshal(body, &u); err != nil {
//		log.Println("Failed to unmarshal payload:", err)
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//}
