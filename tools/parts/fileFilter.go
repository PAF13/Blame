package parts

func (liste *BETRIEBSMITELLLISTE) fileFilter(filter *FILTER) *BETRIEBSMITELLLISTE {
	listeNew := NewBetriebsmillliste()
	listeRest := NewBetriebsmillliste()
	for _, b := range liste.Betriebsmittel {
		for _, bb := range b.Artikel {
			if (bb.Funktionsgruppe != "Klemme" && bb.Hersteller != "HELU" && bb.Hersteller != "LAP") && filter.Filter["=="+b.BMK.FunktionaleZuordnung+"+"+b.BMK.Ortskennzeichen] && bb.Bestellnummer != "" {
				key := "==" + b.BMK.FunktionaleZuordnung + "=" + b.BMK.Funktionskennzeichen + "++" + b.BMK.Aufstellungsort + "+" + b.BMK.Ortskennzeichen + "-" + b.BMK.BMK + "|" + bestellnummerCleaner(bb.Bestellnummer)
				_, ok := listeNew.Betriebsmittel[key]
				if ok {
					listeNew.Betriebsmittel[key].Artikel = append(listeNew.Betriebsmittel[key].Artikel, bb)
				} else {
					listeNew.Betriebsmittel[key] = b.NewBetriebsmittelTemp2()
					listeNew.Betriebsmittel[key].Artikel = append(listeNew.Betriebsmittel[key].Artikel, bb)
				}
			} else {
				key := b.BMK.FunktionaleZuordnung + b.BMK.Funktionskennzeichen + b.BMK.Aufstellungsort + b.BMK.Ortskennzeichen + b.BMK.BMK
				_, ok := listeRest.Betriebsmittel[key]
				if ok {
					listeRest.Betriebsmittel[key].Artikel = append(listeRest.Betriebsmittel[key].Artikel, bb)
				} else {
					listeRest.Betriebsmittel[key] = b.NewBetriebsmittelTemp2()
					listeRest.Betriebsmittel[key].Artikel = append(listeRest.Betriebsmittel[key].Artikel, bb)
				}
			}
		}
	}

	*liste = *listeNew
	return listeRest
}
