package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// https://cloud.google.com/pubsub/docs/push#receive_push
type PubSubMessage struct {
	Message struct {
		Data []byte `json:"data,omitempty"`
		ID   string `json:"id"`
	} `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("io.ReadAll: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Debug
	log.Printf("%s", string(body))

	var m PubSubMessage
	if err := json.Unmarshal(body, &m); err != nil {
		log.Printf("json.Unmarshal: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	name := string(m.Message.Data)
	if name == "" {
		name = "World"
	}
	log.Printf("Hello %s!", name)
}

func main() {
	http.HandleFunc("/", handler)
	port := "8080"
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
