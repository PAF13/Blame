package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Struct definitions
type ConfigurationVariable struct {
	ActiveFormulaEvaluated bool   `json:"activeFormulaEvaluated"`
	Name                   string `json:"name"`
	Value                  string `json:"value"`
}

type TypicalConfigurationData struct {
	ConfigurationVariables []ConfigurationVariable `json:"configurationVariables"`
}

type ConfigurationData struct {
	MacroList                []interface{}            `json:"macroList"`
	TypicalConfigurationData TypicalConfigurationData `json:"typicalConfigurationData"`
}

type Config struct {
	OrganizationName    string            `json:"OrganizationName"`
	Library             string            `json:"library"`
	Configurator        string            `json:"configurator"`
	ConfiguratorElement string            `json:"configuratorElement"`
	ConfigurationData   ConfigurationData `json:"configurationData"`
}

var tom *Config = &Config{
	OrganizationName:    "",
	Library:             "",
	Configurator:        "",
	ConfiguratorElement: "",
	ConfigurationData: ConfigurationData{
		TypicalConfigurationData: TypicalConfigurationData{
			ConfigurationVariables: []ConfigurationVariable{},
		},
	},
}

func tomHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		// Just send out the JSON version of 'tom'
		j, _ := json.Marshal(tom)
		w.Write(j)
	case "POST":
		// Decode the JSON in the body and overwrite 'tom' with it
		d := json.NewDecoder(r.Body)
		p := &Config{}
		err := d.Decode(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		tom = p
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "I can't do that.")
	}
}
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s %s", r.RemoteAddr, r.Method, r.URL.Path, time.Since(start))
	})
}

func main() {

	http.HandleFunc("/", tomHandler)
	mux := http.NewServeMux()
	loggedMux := loggingMiddleware(mux)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", loggedMux); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
