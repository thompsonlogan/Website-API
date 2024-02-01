package main

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
	//Number int64  `json:"number"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type LoginRequest struct {
	//Number   int64  `json:"number"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

/*type TransferRequest struct {
	ToAccount int `json:"toAccount"`
	Amount    int `json:"amount"`
}*/

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

// TODO: remove the balance field from Account struct
type Account struct {
	ID                int       `json:"id"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	Number            int64     `json:"number"`
	EncryptedPassword string    `json:"-"`
	Balance           int64     `json:"balance"`
	Email             string    `json:"email"`
	CreatedAt         time.Time `json:"createdAt"`
}

func (a *Account) ValidPassword(pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.EncryptedPassword), []byte(pw)) == nil
}

func NewAccount(firstName, lastName, password string, email string) (*Account, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Account{
		FirstName:         firstName,
		LastName:          lastName,
		EncryptedPassword: string(encpw),
		Email:             email,
		Number:            int64(rand.Intn(1000000)),
		CreatedAt:         time.Now().UTC(),
	}, nil
}
