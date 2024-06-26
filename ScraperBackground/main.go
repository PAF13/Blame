package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// ConfigurationVariable represents each configuration variable in the JSON
type ConfigurationVariable struct {
	ActiveFormulaEvaluated bool        `json:"activeFormulaEvaluated"`
	Name                   string      `json:"name"`
	Value                  interface{} `json:"value"`
}

// TypicalConfigurationData represents the typical configuration data in the JSON
type TypicalConfigurationData struct {
	ConfigurationVariables []*ConfigurationVariable `json:"configurationVariables"`

	TypicalInstances map[string]interface{} `json:"typicalInstances"`
}

// ConfigurationData represents the configuration data in the JSON
type ConfigurationData struct {
	MacrosList               []interface{}            `json:"macrosList"`
	TypicalConfigurationData TypicalConfigurationData `json:"typicalConfigurationData"`
}

// Message represents the entire JSON payload
type Message struct {
	OrganizationName    string            `json:"OrganizationName"`
	Library             string            `json:"library"`
	Configurator        string            `json:"configurator"`
	ConfiguratorElement string            `json:"configuratorElement"`
	ConfigurationData   ConfigurationData `json:"configurationData"`
}

func logRequest(r *http.Request) {
	// Log the request method and URL
	//log.Printf("%s %s\n", r.Method, r.URL.Path)

	// Log headers
	/*for name, values := range r.Header {
		for _, value := range values {
			log.Printf("Header: %s = %s\n", name, value)
		}
	}*/

	// Log body if the method is POST
	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err == nil {
			var f Message
			json.Unmarshal(body, &f)
			//log.Printf("Body: %s\n", string(body))
			for _, b := range f.ConfigurationData.TypicalConfigurationData.ConfigurationVariables {
				switch b.Value {
				case "false":
					b.Value = false
				case "true":
					b.Value = true

				}
				switch b.Value.(type) {
				case bool:
					fmt.Printf("Variable: %-30s State: %t\n", b.Name, b.Value)
				case string:
					fmt.Printf("Variable: %-30s State: %s\n", b.Name, b.Value)
				}
			}
			//log.Println(f)
			// Restore the body for further processing
			//r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}
	}
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	// Handle CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Log the request details
	logRequest(r)

	// Handle OPTIONS request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "OPTIONS request received"}`))
		return
	}

	if r.Method == http.MethodPost {
		// Read the body of the request
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Unable to read body", http.StatusBadRequest)
			return
		}

		// Unmarshal the JSON into a Message struct
		var msg Message
		err = json.Unmarshal(body, &msg)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// Print the received message to the console
		//fmt.Printf("Received message: %+v\n", msg)

		// Respond with a success message
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Message received"}`))
	} else {
		// Respond with a method not allowed error
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/receive-message", messageHandler)
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
