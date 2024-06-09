package parts

func readEPlanArtikels() {
	dir := "\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Quelldaten\\Eplan2024_Datenbank.xml"

	eplanArtikel := []*ARTIKEL{}
	readXML(dir, eplanArtikel)

}
