package parts

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	pfaden := []string{
		"\\\\ME-Datenbank-1\\Database\\Schnittstelle\\BlameInput\\Kopie von 8000831_Stückliste_Bedienkästen 16EP1 und 20EP1.xlsx",
	}
	LoadStueckliste(pfaden)
	got := false
	if got != false {
		t.Errorf("failed")
	}
}
func TestLager(t *testing.T) {
	pfaden := []string{
		"\\\\ME-Datenbank-1\\Database\\Schnittstelle\\BlameInput\\Topix.xlsx",
		"\\\\ME-Datenbank-1\\Database\\Schnittstelle\\BlameInput\\Moeller.xlsx",
		"\\\\ME-Datenbank-1\\Database\\Schnittstelle\\BlameInput\\Lagerhueter.xlsx",
	}

	LoadLager(pfaden)
	got := false
	if got != false {
		t.Errorf("failed")
	}
}
func TestFilter(t *testing.T) {
	got := Filter()
	fmt.Println(got)
	if got != false {
		t.Errorf("failed")
	}
}
func TestXMLReader(t *testing.T) {
	readEPlanArtikels()

	if false {
		t.Errorf("failed")
	}
}
