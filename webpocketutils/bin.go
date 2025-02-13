package webpocket

import (
	"fmt"
	"flag"
	"log"
	"os"
	"strings"
	"net/http"
	"net/http/httputil"
)

var (
	requestlog string
)

func init() {
	flag.StringVar(&requestlog, "rl", "requestlog.txt", "Output file for request bin. -b needs to be provided")
}

func RequestBin(w http.ResponseWriter, r *http.Request) {

	logFile, err := os.OpenFile(requestlog, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Println("[-] Error creating request bin log file")
	}
	defer logFile.Close()

	req, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("Request from %s:\n%s", r.RemoteAddr, string(req))
	fmt.Fprintf(logFile, "Request from %s:\n%s", r.RemoteAddr, strings.ReplaceAll(string(req), "\r\n", "\n"))
}
