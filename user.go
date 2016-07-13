package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type user struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Gamertag string `json:"gamertag"`
	Bio      string `json:"bio"`
	Mentor   bool   `json:"mentor"`
	Joined   string `json:"joined"`
	Timezone string `json:"timezone"`
	Verified bool   `json:"verified"`
}
type userResults []user

func userHandler(w http.ResponseWriter, r *http.Request) {
	u := user{
		ID:       0,
		Username: "user",
		Gamertag: "gamer123",
		Bio:      "big paragraph follows\nlol\n",
		Mentor:   false,
		Joined:   "2016-01-17 20:43:01 UTC",
		Timezone: "EST",
		Verified: true,
	}

	results := userResults{}
	results = append(results, u)

	response := ocResponse{
		Body: results,
		Error: resError{
			Type:        0,
			Description: "This error was caused by cats",
			Level:       0,
		},
		Meta: meta{
			Records: 1,
			Type:    "user",
			Time:    "2016-01-17 22:57:03 UTC",
		},
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Println(err)
	}

	return
}
