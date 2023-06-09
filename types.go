package main

import (
	"math/rand"
	"time"
)

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	createdAt time.Time `json: "created_at"`
}

type CreateAccountRequest struct {
	FirstName string `json: "first_name"`
	LastName  string `json: "last_name"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{

		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(100000)),
		createdAt: time.Now().UTC(),
	}
}
