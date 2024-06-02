package parts

func Filter() bool {
	var stueckliste BETRIEBSMITELLLISTE
	var filter FILTER
	stuecklisteFiltered := BETRIEBSMITELLLISTE{
		Betriebsmittel: make(map[string]*BETRIEBSMITELL),
	}
	stueckliste.readJson("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Stueckliste\\Stueckliste_Clean.json")
	filter.readJson("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Stueckliste\\filter.json")

	for _, b := range stueckliste.Betriebsmittel {
		_, ok := filter.Filter[b.BMK.Ortskennzeichen]

		if ok && filter.Filter[b.BMK.Ortskennzeichen] {

			for _, bb := range b.Artikel {
				var part *BETRIEBSMITELL
				_, okPart := stuecklisteFiltered.Betriebsmittel[b.BMK.Ortskennzeichen+":"+bb.Bestellnummer]
				if !okPart {
					part = NewBetriebsmittel()
					part.BMK.Ortskennzeichen = b.BMK.Ortskennzeichen
					stuecklisteFiltered.Betriebsmittel[b.BMK.Ortskennzeichen+":"+bb.Bestellnummer] = part
					bb.Bestellung_Siteca = bb.Stueckzahl
					part.Artikel = append(part.Artikel, bb)
				} else {
					part = stuecklisteFiltered.Betriebsmittel[b.BMK.Ortskennzeichen+":"+bb.Bestellnummer]
					part.Artikel[0].Stueckzahl = part.Artikel[0].Stueckzahl + bb.Stueckzahl
					part.Artikel[0].Bestellung_Siteca = part.Artikel[0].Bestellung_Siteca + bb.Stueckzahl
				}
			}
		}
	}
	writeStueckliste("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Stueckliste\\Stueckliste_Filtered", stuecklisteFiltered.Betriebsmittel)
	writeJsonFile("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Stueckliste\\Stueckliste_Filtered", stuecklisteFiltered.Betriebsmittel)
	writeJsonFile("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Stueckliste\\Filtered", filter)
	return false
}
