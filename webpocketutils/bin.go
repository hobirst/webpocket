package webpocket

import (
	"fmt"
	"log"
	"os"
	"strings"
	"net/http"
	"net/http/httputil"
)


func RequestBin(w http.ResponseWriter, r *http.Request) {

	logFile, err := os.OpenFile(requestlog, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		if !Quite {
			log.Println("[-] Error creating request bin log file")
		}
		return
	}
	defer logFile.Close()

	req, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Println(err)
		return
	}

	if !Quite {
		fmt.Printf("Request from %s:\n%s", r.RemoteAddr, string(req))
		fmt.Fprintf(logFile, "Request from %s:\n%s", r.RemoteAddr, strings.ReplaceAll(string(req), "\r\n", "\n"))
	}
}
