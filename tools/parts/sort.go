package parts

func (liste *BETRIEBSMITELLLISTE) lagerstandabgleich(beistellung *BETRIEBSMITELLLISTE) {
	lager := NewLagerliste()
	lager.readJson("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Lager\\Lager.json")

	for a, b := range liste.Betriebsmittel {

		for _, bb := range b.Artikel {
			var key string
			if bb.Bestellnummer != "" {
				key = bestellnummerCleaner(bb.Bestellnummer)
			} else {
				key = bb.ERP_KNT
			}
			_, okBeistellung := beistellung.Betriebsmittel[a]
			if okBeistellung && beistellung.Betriebsmittel[a].Artikel[0].Beigestellt {
				vergleich(bb, &bb.Beistellung_Stueckzahl, &beistellung.Betriebsmittel[a].Artikel[0].Stueckzahl)
			}
			_, ok := lager.Betriebsmittel[key]
			if ok {
				bb.ERP = lager.Betriebsmittel[key].ERP

				vergleich(bb, &bb.Bestellung_Moeller, &lager.Betriebsmittel[key].Bestellung_Moeller)
				vergleich(bb, &bb.Lager_Siteca, &lager.Betriebsmittel[key].Lager_Siteca)
				vergleich(bb, &bb.Bestellung_KNT, &lager.Betriebsmittel[key].Bestellung_KNT)
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
