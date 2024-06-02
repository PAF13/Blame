package parts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func (f *BETRIEBSMITELLLISTE) readJson(pfad string) *BETRIEBSMITELLLISTE {
	steucklisteClean, err := os.Open(pfad)
	if err != nil {
		fmt.Println(err)
	}
	defer steucklisteClean.Close()
	byteValue, _ := ioutil.ReadAll(steucklisteClean)
	json.Unmarshal(byteValue, f)
	return f
}

func (f *FILTER) readJson(pfad string) *FILTER {
	steucklisteClean, err := os.Open(pfad)
	if err != nil {
		fmt.Println(err)
	}
	defer steucklisteClean.Close()
	byteValue, _ := ioutil.ReadAll(steucklisteClean)
	json.Unmarshal(byteValue, f)
	return f
}
