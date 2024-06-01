package parts

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	got := ReadFile("C:\\Dev\\Blame\\tools\\parts\\test\\Files\\240529_8000634-02_Artikelstückliste_DE_Siteca.xlsx")
	fmt.Println(got)
	if got != false {
		t.Errorf("failed")
	}
}
