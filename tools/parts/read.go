package parts

import (
	"fmt"
	"strings"
)

func load(pfad string) (*[][]string, string) {
	fileName := strings.Split(strings.Split(pfad, "\\")[len(strings.Split(pfad, "\\"))-1], ".")[0]
	fileName = "Blame_" + fileName
	rows := readExcel(pfad)
	return rows, fileName
}
func LoadStueckliste(pfaden []string) {

	for _, pfad := range pfaden {
		betriebsmittelListe := NewBetriebsmillliste()
		beistellung := NewBetriebsmillliste()
		filter := NewFilter()
		filter2 := NewFilter()
		rows, fileName := load(pfad)
		rows2, _ := load("\\\\ME-Datenbank-1\\Projektdaten 2024\\KROENERT\\8000634_Kay II-Automotive\\03 MATERIAL\\01 STUECKLISTEN\\40514_8000634-02_Beistellung_Siteca_DE.xlsx")
		header, headerRow, err := setHeader(rows)
		if err != nil {
			fmt.Println(err)
		}
		header2, headerRow2, err := setHeader(rows2)
		if err != nil {
			fmt.Println(err)
		}

		betriebsmittelListe.setListe(rows, header, headerRow, false, filter)
		//betriebsmittelListe.writeJsonFile(fileName + "_Raw")
		//writeStueckliste("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Stueckliste\\"+fileName+"_Raw", betriebsmittelListe.Betriebsmittel)

		beistellung.setListe(rows2, header2, headerRow2, true, filter2)
		//beistellung.writeJsonFile(fileName2 + "_Raw")
		//writeStueckliste("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Stueckliste\\"+fileName2+"_Raw", beistellung.Betriebsmittel)

		//writeJsonFile("Filter", filter)

		betriebsmittelListeRest := betriebsmittelListe.fileFilter(filter)
		//betriebsmittelListe.writeJsonFile(fileName + "_Filter")
		//writeStueckliste("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Stueckliste\\"+fileName+"_Rest", betriebsmittelListeRest.Betriebsmittel)

		betriebsmittelListe.listSum(filter)
		//betriebsmittelListe.writeJsonFile(fileName + "_Sum")
		//writeStueckliste("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Stueckliste\\"+fileName+"_Rest", betriebsmittelListeRest.Betriebsmittel)

		beistellung.listSum(filter)
		//beistellung.writeJsonFile(fileName2 + "_Sum")
		//writeStueckliste("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Stueckliste\\"+fileName+"_Rest", betriebsmittelListeRest.Betriebsmittel)

		betriebsmittelListeRest.listSum2()
		//betriebsmittelListeRest.writeJsonFile(fileName + "_SumRest")
		//writeStueckliste("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Stueckliste\\"+fileName+"_Rest", betriebsmittelListeRest.Betriebsmittel)

		betriebsmittelListe.lagerstandabgleich(beistellung)
		betriebsmittelListe.writeJsonFile(fileName)
		writeStueckliste("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Stueckliste\\"+fileName, betriebsmittelListe.Betriebsmittel)

		betriebsmittelListeRest.lagerstandabgleich(beistellung)
		betriebsmittelListeRest.writeJsonFile(fileName + "_Rest")
		writeStueckliste("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Stueckliste\\"+fileName+"_Rest", betriebsmittelListeRest.Betriebsmittel)

		var orten [][]string

		for a, b := range filter.Filter {
			var ort []string
			ort = append(ort, a)
			if b {
				ort = append(ort, "True")
			} else {
				ort = append(ort, "False")
			}
			orten = append(orten, ort)
		}
		//writeCSVFile("fileName", &orten, "EPlanBetriebsmittel")
	}

}

func LoadLager(pfaden []string) {
	lagerliste := NewLagerliste()
	for _, pfad := range pfaden {

		fileName := strings.Split(strings.Split(pfad, "\\")[len(strings.Split(pfad, "\\"))-1], ".")[0]
		fmt.Println(fileName)
		rows := readExcel(pfad)
		header, headerRow, err := setHeader(rows)
		if err != nil {
			fmt.Println(err)
		}
		lagerliste.setListe(rows, header, headerRow)
	}
	lagerliste.writeJsonFile("Lager")
	lagerliste.writeStueckliste("Lager")
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
