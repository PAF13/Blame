package parts

func (liste *BETRIEBSMITELLLISTE) listSum(filter *FILTER) {
	listeNew := NewBetriebsmillliste()
	for _, b := range liste.Betriebsmittel {

		for _, bb := range b.Artikel {
			if filter.Filter["=="+b.BMK.FunktionaleZuordnung+"+"+b.BMK.Ortskennzeichen] {
				//preKey := "==" + b.BMK.FunktionaleZuordnung + "=" + b.BMK.Funktionskennzeichen + "++" + b.BMK.Aufstellungsort + "+" + b.BMK.Ortskennzeichen + "-" + b.BMK.BMK + "|"
				preKey := "==" + b.BMK.FunktionaleZuordnung + "+" + b.BMK.Ortskennzeichen + "|"
				var key string
				if bb.Bestellnummer != "" {
					key = preKey + bestellnummerCleaner(bb.Bestellnummer)
				} else {
					key = preKey + bb.ERP_KNT
				}
				keyBetriebsmittel := "==" + b.BMK.FunktionaleZuordnung +
					"=" + b.BMK.Funktionskennzeichen +
					"++" + b.BMK.Aufstellungsort +
					"+" + b.BMK.Ortskennzeichen +
					"-" + b.BMK.BMK +
					"|" + bestellnummerCleaner(bb.Bestellnummer)
				_, ok := listeNew.Betriebsmittel[key]
				if ok {
					listeNew.Betriebsmittel[key].Artikel[0].Stueckzahl = listeNew.Betriebsmittel[key].Artikel[0].Stueckzahl + bb.Stueckzahl
					listeNew.Betriebsmittel[key].Artikel[0].Bestellung_Siteca = listeNew.Betriebsmittel[key].Artikel[0].Stueckzahl
				} else {
					listeNew.Betriebsmittel[key] = b.NewBetriebsmittelTemp3()

					listeNew.Betriebsmittel[key].Artikel = append(listeNew.Betriebsmittel[key].Artikel, bb)
					listeNew.Betriebsmittel[key].Artikel[0].Bestellung_Siteca = bb.Stueckzahl
				}
				listeNew.Betriebsmittel[key].SumBauteile[keyBetriebsmittel] = &ARTIKEL{}
			}
		}
	}
	*liste = *listeNew
}
func (liste *BETRIEBSMITELLLISTE) listSum2() {
	listeNew := NewBetriebsmillliste()
	for _, b := range liste.Betriebsmittel {

		for _, bb := range b.Artikel {
			var key string
			if bb.Bestellnummer != "" {
				key = bb.Bestellnummer
			} else {
				key = bb.ERP_KNT
			}
			_, ok := listeNew.Betriebsmittel[key]
			if ok {
				listeNew.Betriebsmittel[key].Artikel[0].Stueckzahl = listeNew.Betriebsmittel[key].Artikel[0].Stueckzahl + bb.Stueckzahl
				listeNew.Betriebsmittel[key].Artikel[0].Bestellung_Siteca = listeNew.Betriebsmittel[key].Artikel[0].Stueckzahl
			} else {
				listeNew.Betriebsmittel[key] = NewBetriebsmittel()
				listeNew.Betriebsmittel[key].Artikel = append(listeNew.Betriebsmittel[key].Artikel, bb)
				listeNew.Betriebsmittel[key].Artikel[0].Bestellung_Siteca = bb.Stueckzahl
			}

		}
	}
	*liste = *listeNew
}
