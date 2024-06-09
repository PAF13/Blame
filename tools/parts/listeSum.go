package parts

import "strings"

func (liste *BETRIEBSMITELLLISTE) listSum() {
	listeNew := NewBetriebsmillliste()
	for a, b := range liste.Betriebsmittel {

		for _, bb := range b.Artikel {

			_, ok := listeNew.Betriebsmittel[a]
			if strings.Contains(bb.Quelle, "Schaltschrank") {
				if ok {
					listeNew.Betriebsmittel[a].Artikel[0].Stueckzahl = listeNew.Betriebsmittel[a].Artikel[0].Stueckzahl + bb.Stueckzahl
					listeNew.Betriebsmittel[a].Artikel[0].Bestellung_Siteca = listeNew.Betriebsmittel[a].Artikel[0].Stueckzahl
				} else {
					listeNew.Betriebsmittel[a] = b.NewBetriebsmittelTemp3()

					listeNew.Betriebsmittel[a].Artikel = append(listeNew.Betriebsmittel[a].Artikel, bb)
					listeNew.Betriebsmittel[a].Artikel[0].Bestellung_Siteca = bb.Stueckzahl
				}
			} else {
				if ok {
					listeNew.Betriebsmittel[a].Artikel[0].Stueckzahl = listeNew.Betriebsmittel[a].Artikel[0].Stueckzahl + bb.Stueckzahl
					listeNew.Betriebsmittel[a].Artikel[0].Bestellung_Siteca = listeNew.Betriebsmittel[a].Artikel[0].Stueckzahl
				} else {
					listeNew.Betriebsmittel[a] = b.NewBetriebsmittelTemp4()

					listeNew.Betriebsmittel[a].Artikel = append(listeNew.Betriebsmittel[a].Artikel, bb)
					listeNew.Betriebsmittel[a].Artikel[0].Bestellung_Siteca = bb.Stueckzahl
				}
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
