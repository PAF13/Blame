package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)



var einstellungen *EINSTELLUNGEN

func INIT_SETTINGS() {
	einstellungen = &EINSTELLUNGEN{}
	einstellungen.ALLEMEIN_EINSTELLUNG = ALLGEMEIN_EINSTELLUNG{
		PROJEKTORDNER: PROJEKTORDNER{
			ROOT: "\\\\ME-Datenbank-1\\Projektdaten ",
			DOKUMENTE: DOKUMENTE{
				ROOT: "\\01 DOKUMENTE",
			},
			SCHALTPLAN: SCHALTPLAN{
				ROOT: "\\02 SCHALTPLAN",
				KUNDENAUSFERTIGUNG: "\\02 KUNDENAUSFERTIGUNG",
			},
			MATERIAL: MATERIAL{
				ROOT: "\\03 MATERIAL",
				STUECKLISTE_KUNDENFREIGABE: "\\02 KUNDENFREIGABE",
				STUECKLISTE_INTERN: "\\03 INTERN",
				STUECKLISTE_TOPIX: "\\04 TOPIX ERLEDIGT",
			},
			FERTIGUNGSDATEN: FERTIGUNGSDATEN{
				ROOT: "\\04 FERTIGUNGSDATEN",
				NC: "\\01 NC",
				DRAHT: "\\02 DRAHT",
				BESCHRIFTUNG: "\\03 BESCHRIFTUNG",
				UNTERLAGEN: "\\04 UNTERLAGEN",
			},
			FOTOS: FOTOS{
				ROOT: "\\05 FOTOS",
				BEISTELLUNG: "\\01 BEISTELLUNG",
				PRODUKTION: "\\02 PRODUKTION",
			},
			PRUEFPROTOKOLL: PRUEFPROTOKOLL{
				ROOT: "\\06 PRUEFPROTOKOLLE KUNDENAUSFERTIGUNG",
			},
		},
		STD_PFAD: STD_PFAD{
			ROOT: "C:\\Dev\\Blame\\frontend\\src\\api",
		},
	}
	einstellungen.KUNDE_EINSTELLUNG = map[string]KUNDE_EINSTELLUNG{}
	einstellungen.KUNDE_EINSTELLUNG["KNT"] = KUNDE_EINSTELLUNG{
		KUNDE_STUECKLISTE: KUNDE_STUECKLISTE{
			KUNDE_BMK: KUNDE_BMK{
				BMK_VOLL:                   9999,
				BMK_ID:                     9999,
				FUNKTIONALEZUORDNUNG:       0,
				FUNKTIONSKENNZEICHEN:       1,
				AUFSTELLUNGSORT:            2,
				ORTSKENNZEICHEN:            3,
				BMK:                        4,
				DOKUMENTENART:              9999,
				BENUTZERDEFINIERTESTRUKTUR: 9999,
				ANLAGENNUMMER:              9999,
				KENNBUCHSTABE:              9999,
			},
			FIRST_VALUE: 9999,
			ARTIKEL: ARTIKEL{
				ERP:                 7,
				ERP_QUELLE:          9999,
				BESTELLNUMMER:       9,
				ARTIKELNUMMER_EPLAN: 9999,
				HERSTELLER:          11,
				STEUCKZAHL:          5,
				EINHEIT:             9999,
				BEISTELLUNG:         12,
				GELIEFERT:           9999,
			},
		},
		KUNDE_LAGERBESTAND: KUNDE_LAGERBESTAND{
			LAGERORT: 9999,
			FIRST_VALUE: 9999,
			ARTIKEL: ARTIKEL{
				ERP:                 9999,
				ERP_QUELLE:          9999,
				BESTELLNUMMER:       9999,
				ARTIKELNUMMER_EPLAN: 9999,
				HERSTELLER:          9999,
				STEUCKZAHL:          9999,
				EINHEIT:             9999,
				BEISTELLUNG:         9999,
				GELIEFERT:           9999,
			},
		},
	}

	einstellungen.KUNDE_EINSTELLUNG["TIG"] = KUNDE_EINSTELLUNG{
		KUNDE_STUECKLISTE: KUNDE_STUECKLISTE{
			KUNDE_BMK: KUNDE_BMK{
				BMK_VOLL:                   9999,
				BMK_ID:                     9999,
				FUNKTIONALEZUORDNUNG:       9999,
				FUNKTIONSKENNZEICHEN:       9999,
				AUFSTELLUNGSORT:            9999,
				ORTSKENNZEICHEN:            9999,
				BMK:                        9999,
				DOKUMENTENART:              9999,
				BENUTZERDEFINIERTESTRUKTUR: 9999,
				ANLAGENNUMMER:              9999,
				KENNBUCHSTABE:              9999,
			},
			FIRST_VALUE: 9999,
			ARTIKEL: ARTIKEL{
				ERP:                 9999,
				ERP_QUELLE:          9999,
				BESTELLNUMMER:       9999,
				ARTIKELNUMMER_EPLAN: 9999,
				HERSTELLER:          9999,
				STEUCKZAHL:          9999,
				EINHEIT:             9999,
				BEISTELLUNG:         9999,
				GELIEFERT:           9999,
			},
		},
		KUNDE_LAGERBESTAND: KUNDE_LAGERBESTAND{
			LAGERORT: 9999,
			FIRST_VALUE: 9999,
			ARTIKEL: ARTIKEL{
				ERP:                 9999,
				ERP_QUELLE:          9999,
				BESTELLNUMMER:       9999,
				ARTIKELNUMMER_EPLAN: 9999,
				HERSTELLER:          9999,
				STEUCKZAHL:          9999,
				EINHEIT:             9999,
				BEISTELLUNG:         9999,
				GELIEFERT:           9999,
			},
		},
	}

	einstellungen.KUNDE_EINSTELLUNG["SITECA"] = KUNDE_EINSTELLUNG{
		KUNDE_STUECKLISTE: KUNDE_STUECKLISTE{
			KUNDE_BMK: KUNDE_BMK{
				BMK_VOLL:                   9999,
				BMK_ID:                     9999,
				FUNKTIONALEZUORDNUNG:       9999,
				FUNKTIONSKENNZEICHEN:       9999,
				AUFSTELLUNGSORT:            9999,
				ORTSKENNZEICHEN:            9999,
				BMK:                        9999,
				DOKUMENTENART:              9999,
				BENUTZERDEFINIERTESTRUKTUR: 9999,
				ANLAGENNUMMER:              9999,
				KENNBUCHSTABE:              9999,
			},
			FIRST_VALUE: 3,
			ARTIKEL: ARTIKEL{
				ERP:                 2,
				ERP_QUELLE:          9999,
				BESTELLNUMMER:       72,
				ARTIKELNUMMER_EPLAN: 187,
				HERSTELLER:          6,
				STEUCKZAHL:          50,
				EINHEIT:             12,
				BEISTELLUNG:         9999,
				GELIEFERT:           9999,
			},
		},
		KUNDE_LAGERBESTAND: KUNDE_LAGERBESTAND{
			LAGERORT: 9999,
			FIRST_VALUE: 9999,
			ARTIKEL: ARTIKEL{
				ERP:                 9999,
				ERP_QUELLE:          9999,
				BESTELLNUMMER:       9999,
				ARTIKELNUMMER_EPLAN: 9999,
				HERSTELLER:          9999,
				STEUCKZAHL:          9999,
				EINHEIT:             9999,
				BEISTELLUNG:         9999,
				GELIEFERT:           9999,
			},
		},
	}

	b2, err := json.MarshalIndent(einstellungen, "", "    ")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(b2))
	err = os.WriteFile("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\Test_Project\\.blame_Einstellung.json", b2, 0644)
	if err != nil {
		log.Println(err)
	}
}