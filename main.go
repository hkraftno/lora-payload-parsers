package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/hkraft/hkraft-iot/lora-payload-parsers/pir_lab_xxns"
	"github.com/hkraft/hkraft-iot/lora-payload-parsers/pul_lab_xxns"
	"github.com/hkraft/hkraft-iot/lora-payload-parsers/tem_lab_xxns"
	"github.com/hkraft/hkraft-iot/lora-payload-parsers/thy_lab_xxns"
	"github.com/hkraft/hkraft-iot/lora-payload-parsers/tor_lab_xxns"
)

func main() {
	http.HandleFunc("/", rootHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Serving http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.URL.String(), "/")
	if r.Method != "GET" {
		http.Error(w, "Only GET is supported", 405)
		return
	} else if len(urlParts) != 3 || urlParts[2] == "" {
		http.Error(w, "Expected HEX to come after /", 400)
		return
	}

	parserName := urlParts[1]
	payloadString := urlParts[2]

	var parser func([]byte) ([]byte, error)
	switch parserName {
	case "pul_lab_xxns":
		parser = pul_lab_xxns.Parse
	case "pir_lab_xxns":
		parser = pir_lab_xxns.Parse
	case "tem_lab_xxns":
		parser = tem_lab_xxns.Parse
	case "thy_lab_xxns":
		parser = thy_lab_xxns.Parse
	case "tor_lab_xxns":
		parser = tor_lab_xxns.Parse
	default:
		http.Error(w, "Unknown parser "+parserName, 404)
		return
	}

	hexBytes, err := hex.DecodeString(payloadString)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	jsonString, err := parser(hexBytes)
	if err != nil {
		message := fmt.Sprintf("Got error parsing they payload %s: %s", payloadString, err.Error())
		http.Error(w, message, 400)
		return
	}
	//panic(fmt.Sprintf("%v", jsonString))
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, "%s", jsonString)
}
