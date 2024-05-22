package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type ArtikelSeed struct {
	Bestellnummer string
	Artikelnummer string
	URL           string
}

func main() {
	//importJSON()
	//Scraper()
	//scraperSiemens()
	XMLEplan()

}

func clean(nummer string) string {
	nummer = strings.ReplaceAll(nummer, " ", "")
	nummer = strings.ReplaceAll(nummer, "\t", "")
	nummer = strings.ReplaceAll(nummer, "\n", "")
	nummer = strings.ReplaceAll(nummer, ".", "")
	nummer = strings.ReplaceAll(nummer, "/", "")
	nummer = strings.ReplaceAll(nummer, ",", "")
	nummer = strings.ReplaceAll(nummer, "ü", "ue")
	nummer = strings.ReplaceAll(nummer, "ä", "ae")
	nummer = strings.ReplaceAll(nummer, "ö", "oe")
	nummer = strings.ReplaceAll(nummer, "Ü", "Ue")
	nummer = strings.ReplaceAll(nummer, "Ä", "Ae")
	nummer = strings.ReplaceAll(nummer, "Ö", "Oe")
	nummer = strings.ReplaceAll(nummer, "_", "")
	nummer = strings.ReplaceAll(nummer, "-", "")
	nummer = strings.ReplaceAll(nummer, ":", "")
	nummer2 := strings.ToUpper(nummer)
	nummer2 = strings.ReplaceAll(nummer2, "A", "")
	nummer2 = strings.ReplaceAll(nummer2, "B", "")
	nummer2 = strings.ReplaceAll(nummer2, "C", "")
	nummer2 = strings.ReplaceAll(nummer2, "D", "")
	nummer2 = strings.ReplaceAll(nummer2, "E", "")
	nummer2 = strings.ReplaceAll(nummer2, "F", "")
	nummer2 = strings.ReplaceAll(nummer2, "G", "")
	nummer2 = strings.ReplaceAll(nummer2, "H", "")
	nummer2 = strings.ReplaceAll(nummer2, "I", "")
	nummer2 = strings.ReplaceAll(nummer2, "J", "")
	nummer2 = strings.ReplaceAll(nummer2, "K", "")
	nummer2 = strings.ReplaceAll(nummer2, "L", "")
	nummer2 = strings.ReplaceAll(nummer2, "M", "")
	nummer2 = strings.ReplaceAll(nummer2, "N", "")
	nummer2 = strings.ReplaceAll(nummer2, "O", "")
	nummer2 = strings.ReplaceAll(nummer2, "P", "")
	nummer2 = strings.ReplaceAll(nummer2, "Q", "")
	nummer2 = strings.ReplaceAll(nummer2, "R", "")
	nummer2 = strings.ReplaceAll(nummer2, "S", "")
	nummer2 = strings.ReplaceAll(nummer2, "T", "")
	nummer2 = strings.ReplaceAll(nummer2, "U", "")
	nummer2 = strings.ReplaceAll(nummer2, "V", "")
	nummer2 = strings.ReplaceAll(nummer2, "W", "")
	nummer2 = strings.ReplaceAll(nummer2, "X", "")
	nummer2 = strings.ReplaceAll(nummer2, "Y", "")
	nummer2 = strings.ReplaceAll(nummer2, "Z", "")
	return nummer2
}

func writeJsonFile(dataStruct any, fileName string) {
	data, err := json.MarshalIndent(dataStruct, "", "\t")
	if err != nil {
		log.Println(err)
	}
	err = os.WriteFile("C:\\Dev\\Blame\\artikelSeeds\\Clean\\Blame_"+fileName+".json", data, 0644)
	if err != nil {
		log.Println(err)
	}
}

func readJson(fileName string, produkte map[string]*Product) {
	jsonFile_LISTE, err := os.Open("C:\\Dev\\Blame\\artikelSeeds\\Raw\\" + fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile_LISTE.Close()

	byteValue_LISTE, _ := ioutil.ReadAll(jsonFile_LISTE)

	json.Unmarshal(byteValue_LISTE, &produkte)
}
