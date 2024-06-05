package parts

func (liste *BETRIEBSMITELLLISTE) lagerstandabgleich(beistellung *BETRIEBSMITELLLISTE) {
	lager := NewLagerliste()
	lager.readJson("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Lager\\Lager.json")

	for a, b := range liste.Betriebsmittel {

		for _, bb := range b.Artikel {
			_, okBeistellung := beistellung.Betriebsmittel[a]
			if okBeistellung && beistellung.Betriebsmittel[a].Artikel[0].Beigestellt {
				vergleich(bb, &bb.Beistellung_Stueckzahl, &beistellung.Betriebsmittel[a].Artikel[0].Stueckzahl)
			}
			nummer := bestellnummerCleaner(bb.Bestellnummer)
			_, ok := lager.Betriebsmittel[nummer]
			if ok {
				bb.ERP = lager.Betriebsmittel[nummer].ERP

				vergleich(bb, &bb.Bestellung_Moeller, &lager.Betriebsmittel[nummer].Bestellung_Moeller)
				vergleich(bb, &bb.Lager_Siteca, &lager.Betriebsmittel[nummer].Lager_Siteca)
				vergleich(bb, &bb.Bestellung_KNT, &lager.Betriebsmittel[nummer].Bestellung_KNT)
			}

		}
	}
}

func vergleich(bb *ARTIKEL, stueckzahl *float64, stueckzahlLager *float64) {
	if bb.Bestellung_Siteca > 0 && bb.Bestellung_Siteca > *stueckzahlLager {
		*stueckzahl = *stueckzahlLager
		bb.Bestellung_Siteca = bb.Bestellung_Siteca - *stueckzahlLager
		*stueckzahlLager = 0
	} else if bb.Bestellung_Siteca > 0 && bb.Bestellung_Siteca <= *stueckzahlLager {
		*stueckzahl = bb.Bestellung_Siteca
		*stueckzahlLager = *stueckzahlLager - bb.Bestellung_Siteca
		bb.Bestellung_Siteca = 0
	}
}
func vergleich2(bb *ARTIKEL, stueckzahl *float64, stueckzahlLager *float64) {
	if bb.Bestellung_Siteca > 0 && bb.Bestellung_Siteca > *stueckzahlLager {
		*stueckzahl = *stueckzahlLager
		bb.Bestellung_Siteca = bb.Bestellung_Siteca - *stueckzahlLager
		*stueckzahlLager = 0
	} else if bb.Bestellung_Siteca > 0 && bb.Bestellung_Siteca <= *stueckzahlLager {
		*stueckzahl = bb.Bestellung_Siteca
		*stueckzahlLager = *stueckzahlLager - bb.Bestellung_Siteca
		bb.Bestellung_Siteca = 0
	}
}
