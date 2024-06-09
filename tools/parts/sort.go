package parts

func (liste *BETRIEBSMITELLLISTE) lagerstandabgleich(beistellung *BETRIEBSMITELLLISTE) {
	lager := NewLagerliste()
	lager.readJson("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Lager\\Lager.json")

	for _, b := range liste.Betriebsmittel {

		for _, bb := range b.Artikel {
			/*_, okBeistellung := beistellung.Betriebsmittel[a]
			if okBeistellung && beistellung.Betriebsmittel[a].Artikel[0].Beigestellt {
				vergleich(bb, &bb.Beistellung_Stueckzahl, &beistellung.Betriebsmittel[a].Artikel[0].Stueckzahl)
			}*/

			nummer := bestellnummerCleaner(bb.Bestellnummer)
			_, ok := lager.Artikel[nummer]
			if ok {
				bb.ERP = lager.Artikel[nummer].ERP

				vergleich(bb, &bb.Bestellung_Moeller, &lager.Artikel[nummer].Bestellung_Moeller)
				vergleich(bb, &bb.Lager_Siteca, &lager.Artikel[nummer].Lager_Siteca)
				vergleich(bb, &bb.Bestellung_KNT, &lager.Artikel[nummer].Bestellung_KNT)
				bb.ARTIKELINFO = lager.Artikel[nummer].ARTIKELINFO
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
