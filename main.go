package main

import (
	"flag"
	"fmt"
	"log"
)

func seedAccount(store Storage, fname, lname, pw string, email string) *Account {
	acc, err := NewAccount(fname, lname, pw, email)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	fmt.Println("new account => ", acc.Number)

	return acc
}

func seedAccounts(s Storage) {
	seedAccount(s, "Logan", "Thompson", "password", "thompsonlogan78@gmail.com")
}

func main() {
	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		fmt.Println("seeding the database")
		seedAccounts(store)
	}

	server := NewAPIServer(":3001", store)
	server.Run()
}
