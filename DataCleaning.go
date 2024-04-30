package main

func Sitecavergleich(artikelstammdaten map[string]*Artikel, liste map[string]*Artikel) map[string]*Artikel {
	for artikel, row := range liste {
		_, ok := artikelstammdaten[artikel]
		if ok {
			liste[artikel] = &Artikel{
				ERP:                artikelstammdaten[artikel].ERP,
				Bestellnummer:      row.Bestellnummer,
				ArtikelnummerEplan: artikelstammdaten[artikel].ArtikelnummerEplan,
				Hersteller:         artikelstammdaten[artikel].Hersteller,
				Beschreibung:       artikelstammdaten[artikel].Beschreibung,
				Stueckzahl:         row.Stueckzahl,
				Warengruppe:        row.Warengruppe,
				Quelle:             "Siteca",
				Stand:              "Topix verhanden",
				Beistellung:        row.Beistellung,
				Ort:                row.Ort,
				Aufstellungsort:    row.Aufstellungsort,
				Ortskennzeichen:    row.Ortskennzeichen,
			}
		} else {
			liste[artikel] = &Artikel{
				ERP:                row.ERP,
				Bestellnummer:      row.Bestellnummer,
				ArtikelnummerEplan: row.ArtikelnummerEplan,
				Hersteller:         row.Hersteller,
				Beschreibung:       row.Beschreibung,
				Stueckzahl:         row.Stueckzahl,
				Warengruppe:        row.Warengruppe,
				Quelle:             row.Quelle,
				Stand:              "Topix nicht verhanden",
				Beistellung:        row.Beistellung,
				Ort:                row.Ort,
				Aufstellungsort:    row.Aufstellungsort,
				Ortskennzeichen:    row.Ortskennzeichen,
			}
		}
	}
	return liste
}
