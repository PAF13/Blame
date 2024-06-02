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
