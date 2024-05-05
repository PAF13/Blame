package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name    string
	Age     int
	Email   string
	Phone   string
	Hobbies []string
}

func (a *App) newProject() {
	p := Person{
		Name:  "John Jones",
		Age:   26,
		Email: "johnjones@email.com",
		Phone: "89910119",
		Hobbies: []string{
			"Swimming",
			"Badminton",
		},
	}

	b, err := json.Marshal(p)
	if err != nil {
		log.Fatalf("Unable to marshal due to %s\n", err)
	}

	fmt.Println(string(b))
}

func (a *App) currentProject() {
}

func (a *App) choseProject() {
}