package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	botToken = "YOUR_BOT_TOKEN"
)

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

type Payload struct {
	User User `json:"user"`
}

func handleTelegramOAuthCallback(w http.ResponseWriter, r *http.Request) {
	hash := r.URL.Query().Get("hash")
	payloadB64 := r.URL.Query().Get("payload")
	payloadBytes, err := base64.StdEncoding.DecodeString(payloadB64)
	if err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	var payload Payload
	err = json.Unmarshal(payloadBytes, &payload)
	if err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	h := hmac.New(sha256.New, []byte(botToken))
	h.Write([]byte(payloadB64))
	checkHash := hex.EncodeToString(h.Sum(nil))

	if hash != checkHash {
		http.Error(w, "Invalid hash", http.StatusBadRequest)
		return
	}

	user := payload.User
	userId := user.Id
	firstName := user.FirstName
	lastName := user.LastName
	username := user.Username
	fmt.Println(firstName, lastName, username, userId)
	// Store user information in your database
	// ...
}

func main() {
	http.HandleFunc("/telegram-oauth-callback", handleTelegramOAuthCallback)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
