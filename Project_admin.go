package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var currentProject *JSON_PARSING
var schranknum int
func (a *App) NewProject2() {
	schranknum = 0
	currentProject = &JSON_PARSING{
		NAME: "projekt",
		VERSION: [3]int{0,0,1},
		BODY: PROJEKT{
			PROJEKT_NUMMER: "8000772",
			PROJEKT_BESCHREIBUNG: "Polifilm",
			BAUJAHR: 2024,
			AKTIV: true,
		},
	}

	b2, err := json.MarshalIndent(currentProject, "", "    ")
	if err != nil {
		log.Println(err)
	}


	fmt.Println(string(b2))
	err = os.WriteFile("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\Test_Project\\.blame.json", b2, 0644)
	if err != nil {
		log.Println(err)
	}
}

func (a *App) CurrentProject() {
	fmt.Println("Current Project: " + currentProject.BODY.PROJEKT_NUMMER+ "_"+currentProject.BODY.PROJEKT_BESCHREIBUNG)
	for _,b := range  currentProject.BODY.PRODUKTE{
		log.Println("Schrank BMK:" + b.BMK.BMK_VOLL)
	}
}

func (a *App) AddProdukt() {
	currentProject.addProdukt2()
}

func (a *JSON_PARSING) addProdukt2() {
	a.BODY.PRODUKTE = append(a.BODY.PRODUKTE, PRODUKTE{
		BMK: BMK2{
			BMK_VOLL: "+EC" + fmt.Sprintf("%d", schranknum),
		},
	})
	schranknum++
}