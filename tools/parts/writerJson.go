package parts

import (
	"encoding/json"
	"log"
	"os"
)

func writeJsonFile(fileName string, dataStruct any) {
	data, err := json.MarshalIndent(dataStruct, "", "\t")
	if err != nil {
		log.Println(err)
	}
	err = os.WriteFile(fileName+".json", data, 0644)
	if err != nil {
		log.Println(err)
	}
}

func (dataStruct *LAGERLISTE) writeJsonFile(name string) {
	data, err := json.MarshalIndent(dataStruct, "", "\t")
	if err != nil {
		log.Println(err)
	}
	err = os.WriteFile("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Lager\\"+name+".json", data, 0644)
	if err != nil {
		log.Println(err)
	}
}
func (dataStruct *BETRIEBSMITELLLISTE) writeJsonFile(name string) {
	data, err := json.MarshalIndent(dataStruct, "", "\t")
	if err != nil {
		log.Println(err)
	}
	err = os.WriteFile("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Stueckliste\\"+name+".json", data, 0644)
	if err != nil {
		log.Println(err)
	}
}
