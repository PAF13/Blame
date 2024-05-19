package main

func newArtikel(url string) *Artikel {
	return &Artikel{
		URL:           url,
		Artikelnummer: "",
		Bestellnummer: "",
	}
}

type Artikel struct {
	Bestellnummer string
	Artikelnummer string
	URL           string
}

func newProduct(url string) *Product {
	return &Product{
		Bestellnummer: "",
		Artikelnummer: "",
		URL:           url,
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
	Bestellnummer       string
	Artikelnummer       string
	URL                 string
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
