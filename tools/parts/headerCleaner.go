package parts

import (
	"math"
	"strings"
)

func headerssClean(headers map[string]uint64) (map[string]uint64, error) {
	var err error
	headersClean := map[string]uint64{
		"FunktionaleZuordnung": math.MaxUint64,
		"Funktionskennzeichen": math.MaxUint64,
		"Aufstellungsort":      math.MaxUint64,
		"Ortskennzeichen":      math.MaxUint64,
		"BMK":                  math.MaxUint64,

		"Bestellnummer":          math.MaxUint64,
		"ERP":                    math.MaxUint64,
		"ERP_KNT":                math.MaxUint64,
		"Hersteller":             math.MaxUint64,
		"Stueckzahl":             math.MaxUint64,
		"Beistellung_Stueckzahl": math.MaxUint64,
		"Bestellung_Moeller":     math.MaxUint64,
		"Lager_Siteca":           math.MaxUint64,
		"Lager_KNT":              math.MaxUint64,
		"Lager_Moeller":          math.MaxUint64,
		"Bestellung_KNT":         math.MaxUint64,
		"Bestellung_Siteca":      math.MaxUint64,
		"Beschreibung":           math.MaxUint64,

		"Beistellung":     math.MaxUint64,
		"Funktionsgruppe": math.MaxUint64,

		"EK_KNT":    math.MaxUint64,
		"EK_Siteca": math.MaxUint64,
	}
	translateStuecklisteKNT := map[string]string{
		"MODUL==":       "FunktionaleZuordnung",
		"FUNKTION=":     "Funktionskennzeichen",
		"AUFSTELLORT++": "Aufstellungsort",
		"EINBAUORT+":    "Ortskennzeichen",
		"BMK-":          "BMK",

		"BESTELLNUMMER": "Bestellnummer",
		"ERP-NUMMER":    "ERP_KNT",
		"HERSTELLER":    "Hersteller",
		"STUECKZAHL":    "Stueckzahl",
		"BEZEICHNUNG":   "Beschreibung",

		"TEILEART":        "Beistellung",
		"FUNKTIONSGRUPPE": "Funktionsgruppe",
	}
	translateLagerTopix := map[string]string{
		"HERSTELLERNUMMER": "Bestellnummer",
		"ARTIKELNUMMER":    "ERP",
		"HERSTELLER":       "Hersteller",
		"MENGEAB(L)":       "Lager_Siteca",
		"BESCHREIBUNG2":    "Beschreibung",
		"EK":               "EK_Siteca",
	}
	translateLagerMoeller := map[string]string{
		"ART.-NR.":     "Bestellnummer",
		"INTERNE-NR.":  "ERP",
		"HERSTELLER":   "Hersteller",
		"ANZAHLME":     "Lager_Moeller",
		"ANZAHLSITECA": "Lager_Siteca",
		"BEZEICHNUNG":  "Beschreibung",
	}
	translateLagerKNT := map[string]string{
		"HERSTELLERARTIKEL": "Bestellnummer",
		"ARTIKLENUMMER":     "ERP_KNT",
		"MENGE":             "Lager_KNT",
		"BEZEICHNUNG2":      "Beschreibung",
		"PREISGLD":          "EK_KNT",
	}
	translate := []map[string]string{
		translateStuecklisteKNT,
		translateLagerTopix,
		translateLagerMoeller,
		translateLagerKNT,
	}
	for _, b := range translate {

		for aa, bb := range headers {
			_, ok := b[aa]

			if ok {
				headersClean[b[aa]] = bb
			}
		}
		if headersClean["Bestellnummer"] != math.MaxUint64 &&
			(headersClean["Lager_Moeller"] != math.MaxUint64 ||
				headersClean["Lager_KNT"] != math.MaxUint64 ||
				headersClean["Stueckzahl"] != math.MaxUint64 ||
				headersClean["Lager_Siteca"] != math.MaxUint64) {
			err = nil
			break
		} else {
			err = New("missing: ")
		}
	}
	/*for a, b := range headersClean {
		fmt.Printf("Header: %-30s value: %d\n", a, b)
	}*/
	return headersClean, err
}

func stringCleaner(x string) string {
	x = strings.ReplaceAll(x, " ", "")
	x = strings.ReplaceAll(x, "\t", "")
	x = strings.ReplaceAll(x, "\n", "")
	x = strings.ToUpper(x)
	x = strings.ReplaceAll(x, "Ü", "UE")
	x = strings.ReplaceAll(x, "Ä", "AE")
	x = strings.ReplaceAll(x, "Ö", "OE")
	return x
}

func bestellnummerCleaner(x string) string {
	x = strings.Replace(x, "TS ", "", 1)
	x = strings.Replace(x, "VX ", "", 1)
	x = strings.Replace(x, "VX.", "", 1)
	x = strings.Replace(x, "SK ", "", 1)
	x = strings.Replace(x, "SV ", "", 1)
	x = strings.Replace(x, "SZ ", "", 1)
	x = strings.Replace(x, "PS ", "", 1)
	x = strings.Replace(x, "EL ", "", 1)
	x = strings.Replace(x, "DK ", "", 1)
	x = strings.Replace(x, "KX ", "", 1)
	x = strings.Replace(x, "TP ", "", 1)

	x = strings.ReplaceAll(x, " ", "")
	x = strings.ReplaceAll(x, "\t", "")
	x = strings.ReplaceAll(x, "\n", "")
	x = strings.ReplaceAll(x, "-", "")
	x = strings.ReplaceAll(x, ".", "")
	x = strings.ReplaceAll(x, "_", "")
	x = strings.ToUpper(x)
	x = strings.ReplaceAll(x, "Ü", "UE")
	x = strings.ReplaceAll(x, "Ä", "AE")
	x = strings.ReplaceAll(x, "Ö", "OE")
	return x
}

func bestellnummerCleaner2(x string) string {
	x = strings.Replace(x, "TS ", "", 1)
	x = strings.Replace(x, "VX ", "", 1)
	x = strings.Replace(x, "VX.", "", 1)
	x = strings.Replace(x, "SK ", "", 1)
	x = strings.Replace(x, "SV ", "", 1)
	x = strings.Replace(x, "SZ ", "", 1)
	x = strings.Replace(x, "PS ", "", 1)
	x = strings.Replace(x, "EL ", "", 1)
	x = strings.Replace(x, "DK ", "", 1)
	x = strings.Replace(x, "KX ", "", 1)
	x = strings.Replace(x, "TP ", "", 1)

	x = strings.ReplaceAll(x, " ", "")
	x = strings.ReplaceAll(x, "\t", "")
	x = strings.ReplaceAll(x, "\n", "")
	x = strings.ReplaceAll(x, ".", "")
	x = strings.ReplaceAll(x, "_", "")
	x = strings.ToUpper(x)
	x = strings.ReplaceAll(x, "Ü", "UE")
	x = strings.ReplaceAll(x, "Ä", "AE")
	x = strings.ReplaceAll(x, "Ö", "OE")
	return x
}

func bestellnummerCleaner3(x string) string {
	x = strings.ReplaceAll(x, " ", "")
	x = strings.ReplaceAll(x, "\t", "")
	x = strings.ReplaceAll(x, "\n", "")
	x = strings.ReplaceAll(x, ".", "")
	x = strings.ReplaceAll(x, "_", "")
	x = strings.ToUpper(x)
	x = strings.ReplaceAll(x, "Ü", "UE")
	x = strings.ReplaceAll(x, "Ä", "AE")
	x = strings.ReplaceAll(x, "Ö", "OE")
	return x
}
