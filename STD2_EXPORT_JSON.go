package main

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
	err = os.WriteFile(rootPfadOutput+"Blame_"+fileName+".json", data, 0644)
	if err != nil {
		log.Println(err)
	}
}
