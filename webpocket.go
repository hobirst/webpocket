package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	webpocket "github.com/hobirst/webpocket/webpocketutils"
)

var (
	helpFlag   bool
	cookies    bool
	requestBin bool
	address    string
	port       string
)

func init() {

	flag.BoolVar(&helpFlag, "h", false, "Print this message")
	flag.BoolVar(&cookies, "c", false, "activate cookiestealer")
	flag.BoolVar(&requestBin, "b", false, "activate request bin")
	flag.StringVar(&address, "a", "0.0.0.0", "Address to listen on")
	flag.StringVar(&port, "p", "6969", "Port\n-p 1234")
	flag.Parse()

}

func main() {

	// print help if wanted
	if helpFlag {
		fmt.Fprintf(os.Stderr, "[-] Usage: %s \n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}

	// init foo
	parserFloat, parserUnit := webpocket.CalcBufferSize()

	// info messages
	log.Printf("[i] Listening on address: %s\n", address)
	log.Printf("[i] Running on port %s\n", port)
	log.Printf("[i] Max upload size: %.2F %s\n", parserFloat, parserUnit)
	if webpocket.Killswitch {
		log.Println("[i] Killswitch activated")
	}

	// server foo
	http.HandleFunc("/f", webpocket.UploadHandler)

	if cookies {
		log.Printf("[i] Cookiestealer activated")
		http.HandleFunc("/c", webpocket.Cookies)
	}

	if requestBin {
		log.Printf("[i] Request bin activated")
		http.HandleFunc("/b", webpocket.RequestBin)
	}

	http.ListenAndServe(address+":"+port, nil)
}
