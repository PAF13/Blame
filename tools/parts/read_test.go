package parts

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	pfaden := []string{
		"C:\\Dev\\Blame\\tools\\parts\\test\\Files\\240529_8000634-02_Artikelstückliste_DE_Siteca.xlsx",
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
