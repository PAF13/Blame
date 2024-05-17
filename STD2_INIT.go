package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"
)

var rootPfad string
var rootPfadOutput string
var rootPfadInput string
var rootPfadDatenbank string
var wg sync.WaitGroup
var pfaden [][]string

func (a *App) BlameStartup() bool {
	start := time.Now()
	rootPfad = "\\\\ME-Datenbank-1\\Database\\Schnittstelle\\"
	rootPfadOutput = rootPfad + "BlameOutput\\"
	rootPfadInput = rootPfad + "BlameInput\\"
	rootPfadDatenbank = rootPfad + "BlameDatenbank\\"
	// Code to measure
	pfaden = [][]string{
		//{"\\\\ME-Datenbank-1\\Database\\Schnittstelle\\BlameInput\\Lagerhueter.xlsx", "KNT", "Lager"},
		//{"\\\\ME-Datenbank-1\\Database\\Schnittstelle\\BlameInput\\Topix.xlsx", "SITECA", "Lager"},
		//{"\\\\ME-Datenbank-1\\Database\\Schnittstelle\\BlameInput\\Moeller.xlsx", "MOELLER", "Lager"},
	}
	for _, pfad := range pfaden {

		pfadLen := len(strings.Split(pfad[0], "\\"))
		fileNameVoll := strings.Split(pfad[0], "\\")[pfadLen-1]
		fileName := strings.Split(fileNameVoll, ".")[0]
		fileExtension := strings.Split(fileNameVoll, ".")[1]
		fmt.Println(fileNameVoll)
		fmt.Println(fileName)
		fmt.Println(fileExtension)

		wg.Add(1)
		ImportFile(pfad[0], pfad[1], pfad[2], fileName)
		wg.Add(1)
		loadFile(pfad[1], pfad[2], fileName)
	}

	duration := time.Since(start)
	fmt.Println(duration)
	fmt.Println(duration.Nanoseconds())
	return true
}

func (a *App) LoadStueckliste(pfad []string, kunde string, fileType string) []string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()
	start := time.Now()
	var temp []string
	for _, pfad2 := range pfad {
		pfadLen := len(strings.Split(pfad2, "\\"))
		fileNameVoll := strings.Split(pfad2, "\\")[pfadLen-1]
		fileName := strings.Split(fileNameVoll, ".")[0]
		fileExtension := strings.Split(fileNameVoll, ".")[1]
		fmt.Println(fileNameVoll)
		fmt.Println(fileName)
		fmt.Println(fileExtension)
		temp = append(temp, fileName)
		wg.Add(1)
		fmt.Println("Importing " + fileName)
		ImportFile(pfad2, kunde, fileType, fileName)
		wg.Add(1)
		fmt.Println("Loading " + fileName)
		loadFile(kunde, fileType, fileName)
		wg.Add(1)
		fmt.Println("Summing " + fileName)
		temp2 := sumListe(kunde, fileType, fileName)

		for _, b := range temp2 {
			temp = append(temp, b)
		}

	}
	wg.Wait()
	fmt.Println("temp")
	fmt.Println(temp)
	duration := time.Since(start)
	fmt.Println(duration)
	fmt.Println(duration.Nanoseconds())
	return temp
}
func (a *App) ExportStueckliste(orten []string, kunde string, fileType string) {
	fmt.Println("trying")
	fileName := orten[0]
	for _, b := range orten {
		fmt.Println(b)
	}
	jsonFile_LISTE, err := os.Open(rootPfadOutput + "Blame_Import3_" + kunde + "_" + fileName + "_" + fileType + ".json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile_LISTE.Close()
	byteValue_LISTE, _ := ioutil.ReadAll(jsonFile_LISTE)
	var artikel_LISTE ARTIKELLISTE
	json.Unmarshal(byteValue_LISTE, &artikel_LISTE)

	writeStueckliste(artikel_LISTE.Artikel, kunde, fileType, fileName)
	writeCSV(&artikel_LISTE, fileName, orten)
}

var currentProject *PROJEKT

func (a *App) NewProject2() {

	currentProject = &PROJEKT{
		PROJEKT_NUMMER:       "8000772",
		PROJEKT_BESCHREIBUNG: "Polifilm",
		BAUJAHR:              2024,
		AKTIV:                true,
	}

	writeJsonFile(".blame", currentProject)
}
