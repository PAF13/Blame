package parts

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
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

		betriebsmittelListe.ProduktDef.FunktionaleZuordnung = true        //==
		betriebsmittelListe.ProduktDef.Funktionskennzeichen = false       //=
		betriebsmittelListe.ProduktDef.Aufstellungsort = false            //++
		betriebsmittelListe.ProduktDef.Ortskennzeichen = true             //+
		betriebsmittelListe.ProduktDef.Dokumentenart = false              //&
		betriebsmittelListe.ProduktDef.BenutzerdefinierteStruktur = false //#
		betriebsmittelListe.ProduktDef.Anlagennummer = false              //empty?
		betriebsmittelListe.ProduktDef.BMK = false                        //-

		beistellung := NewBetriebsmillliste()
		rows, fileName := load(pfad)
		header, headerRow, err := setHeader(rows)
		if err != nil {
			fmt.Println(err)
		}

		betriebsmittelListe.setListe(rows, header, headerRow)
		//betriebsmittelListe.writeJsonFile(fileName + "_Raw")
		//writeStueckliste("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Stueckliste\\"+fileName+"_Raw", betriebsmittelListe.Betriebsmittel)

		//beistellung.setListe(rows2, header2, headerRow2)
		//beistellung.writeJsonFile(fileName2 + "_Raw")
		//writeStueckliste("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Stueckliste\\"+fileName2+"_Raw", beistellung.Betriebsmittel)

		//betriebsmittelListeRest := betriebsmittelListe.fileFilter(filter)
		//betriebsmittelListe.writeJsonFile(fileName + "_Filter")
		//writeStueckliste("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Stueckliste\\"+fileName+"_Rest", betriebsmittelListeRest.Betriebsmittel)

		betriebsmittelListe.listSum()
		//betriebsmittelListe.writeJsonFile(fileName + "_Sum")
		//writeStueckliste("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Stueckliste\\"+fileName+"_Rest", betriebsmittelListeRest.Betriebsmittel)

		//beistellung.listSum(filter2)
		//beistellung.writeJsonFile(fileName2 + "_Sum")
		//writeStueckliste("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Stueckliste\\"+fileName+"_Rest", betriebsmittelListeRest.Betriebsmittel)

		//betriebsmittelListeRest.listSum2()
		//betriebsmittelListeRest.writeJsonFile(fileName + "_SumRest")
		//writeStueckliste("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Stueckliste\\"+fileName+"_Rest", betriebsmittelListeRest.Betriebsmittel)

		betriebsmittelListe.lagerstandabgleich(beistellung)
		betriebsmittelListe.writeJsonFile(fileName)
		writeStueckliste("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Stueckliste\\"+fileName, betriebsmittelListe.Betriebsmittel)

		//betriebsmittelListeRest.lagerstandabgleich(beistellung)
		//betriebsmittelListeRest.writeJsonFile(fileName + "_Rest")
		//writeStueckliste("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Stueckliste\\"+fileName+"_Rest", betriebsmittelListeRest.Betriebsmittel)
	}

}

func LoadLager(pfaden []string) {
	lagerliste := NewLagerliste()
	eplanArtikel := []*ARTIKEL{}
	for _, pfad := range pfaden {

		fileName := strings.ToUpper(strings.Split(strings.Split(pfad, "\\")[len(strings.Split(pfad, "\\"))-1], ".")[0])
		fmt.Println(fileName)
		rows := readExcel(pfad)
		header, headerRow, err := setHeader(rows)
		if err != nil {
			fmt.Println(err)
		}
		lagerliste.setListe(rows, header, headerRow, fileName)
	}

	dir := "\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Quelldaten\\Eplan2024_Datenbank.xml"

	readXML(dir, eplanArtikel)

	//lagerliste.setListe2(artikel)

	lagerliste.writeJsonFile("Lager")
	lagerliste.writeStueckliste("Lager")
}
func readXML(dir string, artikel []*ARTIKEL) {
	xmlFile, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var partsManagement PartsManagement

	xml.Unmarshal(byteValue, &partsManagement)

	eplangruppe := EPLANGRUPPEN{}
	eplangruppe.setInfo()

	for _, b := range partsManagement.Parts {
		artikel = append(artikel, &ARTIKEL{
			Bestellnummer:      b.PArticlePartNr,
			ArtikelnummerEplan: b.PArticleOrderNr,
			ERP:                b.PArticleErpNr,
		})
	}
}

func cleanEplan(partNumber string) string {
	remove := strings.Split(partNumber, ".")
	x := strings.Replace(partNumber, remove[0], "", 1)
	x = bestellnummerCleaner2(x)
	return x
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
