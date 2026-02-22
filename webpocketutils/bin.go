package webpocket

import (
	"fmt"
	"log"
	"os"
	"strings"
	"net/http"
	"net/http/httputil"
)

type RequestBin struct {
	LogFile *os.File
}

func (rb *RequestBin) CreateRequestBinLog(path string) {
	rb.LogFile = CreateLogFile(path)
}

func (rb *RequestBin) ReceiveRequests(w http.ResponseWriter, r *http.Request) {

	req, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Printf("%s Error dumping request: %+v\n", LogErr, err)
		return
	}

	if !Quite {
		fmt.Printf("%s Request from %s:\n%s\n", LogInfo, r.RemoteAddr, string(req))
	}

	fmt.Fprintf(rb.LogFile, "Request from %s:\n%s\n", r.RemoteAddr, strings.ReplaceAll(string(req), "\r\n", "\n"))
	if r.Method == "POST" { // for better readability
		fmt.Fprintf(rb.LogFile, "\n")
	}
}
