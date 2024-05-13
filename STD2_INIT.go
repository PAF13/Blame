package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var rootPfad string
var rootPfadOutput string
var rootPfadInput string
var rootPfadDatenbank string
var wg sync.WaitGroup

func (a *App) BlameStartup() bool {
	start := time.Now()
	// Code to measure

	rootPfad = "\\\\ME-Datenbank-1\\Database\\Schnittstelle\\"
	rootPfadOutput = rootPfad + "BlameOutput\\"
	rootPfadInput = rootPfad + "BlameInput\\"
	rootPfadDatenbank = rootPfad + "BlameDatenbank\\"
	fmt.Println(rootPfadInput)
	fmt.Println(rootPfadDatenbank)
	/*
		lagerSiteca := "Topix_Artikel20240502"
		lagerKNT := "Kopie von Lagerhueter_26_04_2024"

		wg.Add(1)
		go ImportFile2(lagerSiteca, "SITECA", "lager")
		wg.Add(1)
		go ImportFile2(lagerKNT, "KNT", "lager")
		wg.Wait()

		wg.Add(1)
		go loadFile2(lagerSiteca)
		wg.Add(1)
		go loadFile2(lagerKNT)
		wg.Wait()
	*/
	INIT_VERBINDUNGSLITE()
	duration := time.Since(start)
	// Formatted string, such as "2h3m0.5s" or "4.503μs"
	fmt.Println(duration)

	// Nanoseconds as int64
	fmt.Println(duration.Nanoseconds())
	return true
}

func (a *App) LoadStueckliste(pfad string) {
	start := time.Now()
	wg.Add(1)
	ImportFile2(pfad, "stueckliste", "stueckliste")
	wg.Add(1)
	loadFile2("stueckliste")
	wg.Add(1)
	sumListe("stueckliste")
	wg.Wait()

	duration := time.Since(start)
	// Formatted string, such as "2h3m0.5s" or "4.503μs"
	fmt.Println(duration)

	// Nanoseconds as int64
	fmt.Println(duration.Nanoseconds())
}

var currentProject *PROJEKT

func (a *App) NewProject2() {

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
