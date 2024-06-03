package parts

import (
	"fmt"
	"log"
	"strings"

	"github.com/xuri/excelize/v2"
)

func LoadStueckliste(pfaden []string) {
	betriebsmittelListe := NewBetriebsmillliste()

	for _, pfad := range pfaden {
		betriebsmittelListe.importListe(pfad)

	}
	writeStueckliste("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Stueckliste\\Stueckliste", betriebsmittelListe.Betriebsmittel)
	betriebsmittelListe.writeJsonFile("Stueckliste")
}

func LoadLager(pfaden []string) {
	lagerliste := NewLagerliste()
	for _, pfad := range pfaden {
		lagerliste.importListe(pfad)
	}
	lagerliste.writeJsonFile("Lager")
	lagerliste.writeStueckliste("Lager")
}

func (Liste *LAGERLISTE) importListe(pfad string) {
	fileName := strings.Split(strings.Split(pfad, "\\")[len(strings.Split(pfad, "\\"))-1], ".")[0]
	fmt.Println(fileName)
	rows := readExcel(pfad)
	header, headerRow, err := setHeader(rows)
	if err != nil {
		fmt.Println(err)
	}
	Liste.setListe(rows, header, headerRow)

}

func (Liste *BETRIEBSMITELLLISTE) importListe(pfad string) {
	fileName := strings.Split(strings.Split(pfad, "\\")[len(strings.Split(pfad, "\\"))-1], ".")[0]
	fmt.Println(fileName)
	rows := readExcel(pfad)
	header, headerRow, err := setHeader(rows)
	if err != nil {
		fmt.Println(err)
	}
	Liste.setListe(rows, header, headerRow)
}
func readExcel(pfad string) *[][]string {
	file, err := excelize.OpenFile(pfad)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	rows, err := file.GetRows(file.GetSheetList()[0])
	if err != nil {
		log.Fatal(err)
	}
	var maxLength int
	for _, b := range rows {
		if len(b) > maxLength {
			maxLength = len(b)
		}
	}
	fmt.Println(maxLength)
	var rowsClean [][]string
	for _, b := range rows {
		row := make([]string, maxLength)
		copy(row, b)
		rowsClean = append(rowsClean, row)
		//fmt.Println(b)
	}

	return &rowsClean
}

func setHeader(P_rows *[][]string) (map[string]uint64, int, error) {
	var headerRow int
	rows := *P_rows

	headersClean := make(map[string]uint64)
	var err error
	for i := 0; i < 20; i++ {
		headers := make(map[string]uint64)
		for a, b := range rows[i] {
			headers[stringCleaner(b)] = uint64(a)

		}
		headerRow = i
		fmt.Println(headerRow)
		headersClean, err = headerssClean(headers)

		if err != nil {
			fmt.Println(err)
		} else {
			break
		}
	}

	return headersClean, headerRow, err
}
