package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)


func (a *App) BlameStartup() bool {
	//INIT_SETTINGS()
	//INIT_ARTIKELSTAMMDATEN()
	//INIT_VERBINDUNGSLITE()
	//INIT_STUECKLISTE()
	ImportFile()
	return true
}

var currentProject *PROJEKT
var schranknum int

func (a *App) NewProject2() {
	schranknum = 0
	currentProject = &PROJEKT{
		PROJEKT_NUMMER:       "8000772",
		PROJEKT_BESCHREIBUNG: "Polifilm",
		BAUJAHR:              2024,
		AKTIV:                true,
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
	fmt.Println("Current Project: " + currentProject.PROJEKT_NUMMER + "_" + currentProject.PROJEKT_BESCHREIBUNG)
	for _, b := range currentProject.PRODUKTE {
		log.Println("Schrank BMK:" + b.BMK.BMK_VOLL)
	}
}

func (a *App) AddProdukt() {
	currentProject.addProdukt2()
}

func (a *PROJEKT) addProdukt2() {
	a.PRODUKTE = append(a.PRODUKTE, PRODUKTE{
		BMK: BMK2{
			BMK_VOLL: "+EC" + fmt.Sprintf("%d", schranknum),
		},
	})
	schranknum++
}