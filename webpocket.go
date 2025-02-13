package main

import (
	"crypto/tls"
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
	tlsOn      bool
	address    string
	port       string
	certPath   string
	keyPath    string
)

func init() {

	flag.BoolVar(&helpFlag, "h", false, "Print this message")
	flag.BoolVar(&cookies, "c", false, "activate cookiestealer")
	flag.BoolVar(&requestBin, "b", false, "activate request bin")
	flag.BoolVar(&tlsOn, "tls", false, "Activate tls")
	flag.StringVar(&certPath, "cert", "./cert.pem", "Path to certificate")
	flag.StringVar(&keyPath, "key", "./key.pem", "Path to key for certificate")
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

	http.HandleFunc("/f", webpocket.UploadHandler)

	if cookies {
		log.Printf("[i] Cookiestealer activated")
		http.HandleFunc("/c", webpocket.Cookies)
	}

	if requestBin {
		log.Printf("[i] Request bin activated")
		http.HandleFunc("/b", webpocket.RequestBin)
	}


	if tlsOn {
		serverCert, err := tls.LoadX509KeyPair(certPath, keyPath)
		if err != nil {
			log.Fatalf("Error loading cert: %+v\n", err)
		}

		tlsConfig := &tls.Config{
			Certificates: []tls.Certificate{serverCert},
		}

		server := http.Server{
			Addr: address + ":" + port,
			TLSConfig: tlsConfig,
		}
		defer server.Close()

		server.ListenAndServeTLS("", "")
	} else {
		http.ListenAndServe(address+":"+port, nil)
	}



}
