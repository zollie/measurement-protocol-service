package main

import (
	// "bytes"
	"encoding/json"
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/ajg/form"
	"net/http"
	"os"
)

// Sink for incoming measurements
type putFunc func(mp *MeasurementProtocol) error

// Registry for put funcs
var putFuncRegistry = map[string]putFunc{
	"console": putToConsole,
}

// The instance of the Put func
var put putFunc

// Default put to stnd out
func putToConsole(mp *MeasurementProtocol) error {
	fmt.Printf("%#v", mp)
	bytes, err := json.Marshal(mp)
	if err != nil {
		log.Debugf("Error marshalling MP to JSON for debug %s", err.Error())
	} else {
		log.Debugf("\n\nJSON is %s\n\n", bytes)
	}

	return nil
}

// HTTP handler
func handler(w http.ResponseWriter, r *http.Request) {
	var mp *MeasurementProtocol

	log.Debug("Handling new measurement request")

	// Decodes form to struct
	d := form.NewDecoder(r.Body)
	if err := d.Decode(&mp); err != nil {
		log.Debugf("Failed to decode form %s", err.Error())
		http.Error(w, "Form could not be decoded", http.StatusBadRequest)
		return
	}

	log.Debugf("Measurement decoded is: %#v", mp)

	// Puts struct somewhere with Put func
	if err := mp.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	put(mp) // the put
}

// Entrypoint
func main() {
	port := flag.String("port", "2020", "http server listen port")
	logLevel := flag.String("log", "debug", "error|warn|info|debug")
	sink := flag.String("sink", "console", "strategy for storing measurements")
	sinkArgsJson := flag.String("sinkArgs", "{}", "additional args for sink in JSON")

	flag.Parse()

	log.SetOutput(os.Stdout)

	level, err := log.ParseLevel(*logLevel)
	if log.SetLevel(level); err != nil {
		log.Fatalf("Invalid log level %s", logLevel)
	}

	var sinkArgs map[string]*json.RawMessage
	err = json.Unmarshal([]byte(*sinkArgsJson), &sinkArgs)
	if err != nil {
		log.Fatal("Error parsing sink args", err)
	}

	log.Infof("Starting Measurement Protocol service with port %s and sink %s", *port, *sink)

	// Setting the put func
	put = putFuncRegistry[*sink]

	http.HandleFunc("/collect", handler)
	http.ListenAndServe(":"+*port, nil)
}
