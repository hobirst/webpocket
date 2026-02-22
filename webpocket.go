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

func main() {

	if webpocket.HelpFlag {
		fmt.Fprintf(os.Stderr, "%s Usage: %s \n", webpocket.LogErr, os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}

	if webpocket.Killswitch && !webpocket.Quite {
		log.Println("%s Killswitch activated", webpocket.LogInfo)
	}

	http.HandleFunc("/f", webpocket.UploadHandler)

	if webpocket.RunCookieStealer {
		if !webpocket.Quite {
			log.Printf("%s Cookiestealer activated", webpocket.LogInfo)
		}
		cr := &webpocket.CookieReceiver{}
		cr.CreateCookieLog(webpocket.CookieLogPath)
		http.HandleFunc("/c", cr.ReceiveCookies)
	}

	if webpocket.RunRequestBin {
		if !webpocket.Quite {
			log.Printf("%s Request bin activated", webpocket.LogInfo)
		}
		http.HandleFunc("/b/", webpocket.RequestBin)
	}

	if webpocket.RunFileServer {
		//fs := http.FileServer(http.Dir(webpocket.FileServerPath))
		//http.Handle("/fs/", fs)
		http.Handle("/fs/", http.StripPrefix("/fs/", http.FileServer(http.Dir(webpocket.FileServerPath))))
	}

	if webpocket.TlsOn {
		serverCert, err := tls.LoadX509KeyPair(webpocket.CertPath, webpocket.KeyPath)
		if err != nil {
			if !webpocket.Quite {
				log.Fatalf("%s Error loading cert: %+v\n", webpocket.LogErr, err)
			} else {
				return
			}
		}

		tlsConfig := &tls.Config{
			Certificates: []tls.Certificate{serverCert},
		}

		server := http.Server{
			Addr:      webpocket.Address + ":" + webpocket.Port,
			TLSConfig: tlsConfig,
		}
		defer server.Close()

		server.ListenAndServeTLS("", "")
	} else {
		http.ListenAndServe(webpocket.Address + ":" + webpocket.Port, nil)
	}

}
