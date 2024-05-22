package main

func newArtikel(url string) *Artikel {
	return &Artikel{
		URL:           url,
		Artikelnummer: "",
		Bestellnummer: "",
	}
}

type Artikel struct {
	ERPSITECA     string
	ERPKNT        string
	Bestellnummer string
	Artikelnummer string
	EPlannummer   string
	URL           string
	Hersteller    Hersteller
	Anschlusse    []Anschluss
}
type Hersteller struct {
	Hersteller     string
	HerstellerKurz string
	HerstellerLang string
}
type Anschluss struct {
	Name                 string
	Beschreibung         string
	MAXQuerschnitt       float64
	MINQuerschnitt       float64
	Symbol               string
	Technischekenngrösse string
}

func newProductEplan1(e SO117) *Product {
	return &Product{
		ERPTopix:      "",
		Hersteller:    e.P22222,
		Bestellnummer: e.P22002,
		Artikelnummer: e.P22003,
		EAN:           "",
		Kurztext:      "",
		URLHersteller: "",
		URLConrad:     "",
		Generated:     false,
		Lieferinformationen: Lieferinformationen{
			PreisProVE:     0,
			Verfuegbarkeit: "",
			Lieferzeit:     0,
		},
		Produktinformation: Produktinformation{
			Material:     "",
			Oberflaeche:  "",
			Farbe:        "",
			Lieferumfang: "",
			Abmessung: Abmessung{
				Breite: 0.0,
				Hoehe:  0.0,
				Tiefe:  0.0,
			},
			Verpackungsmaße: Abmessung{
				Breite: 0.0,
				Hoehe:  0.0,
				Tiefe:  0.0,
			},
			Schutzarten: "",
			Verpackungseinheit: Verpackungseinheit{
				Anzahl:  0,
				Einheit: "",
			},
			GewichtNetto:  0.0,
			GewichtBrotto: 0.0,
			Beschreibung:  "",
			Ausschreibung: "",
			Ersatzteile:   []string{},
		},
	}
}

func newProductTopix1(e map[string][]string, index int) *Product {
	return &Product{
		ERPTopix:      e["Artikelnummer"][index],
		Hersteller:    e["Hersteller"][index],
		Bestellnummer: e["HerstellerNummer"][index],
		Artikelnummer: e["EPlan_Artikelnr_"][index],
		EAN:           "",
		Kurztext:      "",
		URLHersteller: "",
		URLConrad:     "",
		Generated:     false,
		Lieferinformationen: Lieferinformationen{
			PreisProVE:     0,
			Verfuegbarkeit: "",
			Lieferzeit:     0,
		},
		Produktinformation: Produktinformation{
			Material:     "",
			Oberflaeche:  "",
			Farbe:        "",
			Lieferumfang: "",
			Abmessung: Abmessung{
				Breite: 0.0,
				Hoehe:  0.0,
				Tiefe:  0.0,
			},
			Verpackungsmaße: Abmessung{
				Breite: 0.0,
				Hoehe:  0.0,
				Tiefe:  0.0,
			},
			Schutzarten: "",
			Verpackungseinheit: Verpackungseinheit{
				Anzahl:  0,
				Einheit: "",
			},
			GewichtNetto:  0.0,
			GewichtBrotto: 0.0,
			Beschreibung:  "",
			Ausschreibung: "",
			Ersatzteile:   []string{},
		},
	}
}

type Product struct {
	ERPTopix            string
	Hersteller          string
	Bestellnummer       string
	Artikelnummer       string
	EAN                 string
	Kurztext            string
	URLHersteller       string
	URLConrad           string
	Produktfamilie      string
	ProduktStatus       string
	Generated           bool
	Lieferinformationen Lieferinformationen
	Produktinformation  Produktinformation
}

type Lieferinformationen struct {
	PreisProVE     int
	Verfuegbarkeit string
	Lieferzeit     int
}

type Produktinformation struct {
	Material           string
	Oberflaeche        string
	Farbe              string
	Lieferumfang       string
	Abmessung          Abmessung
	Verpackungsmaße    Abmessung
	Schutzarten        string
	Verpackungseinheit Verpackungseinheit
	GewichtNetto       float64
	GewichtBrotto      float64
	Beschreibung       string
	Ausschreibung      string
	Ersatzteile        []string
}

type Abmessung struct {
	Breite float64
	Hoehe  float64
	Tiefe  float64
}

type Verpackungseinheit struct {
	Anzahl  int
	Einheit string
}
