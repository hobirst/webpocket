package webpocket

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	// flags
	HelpFlag         bool
	RunCookieStealer bool
	RunRequestBin    bool
	TlsOn            bool
	RunFileServer    bool

	ParserSize int

	Address        string
	Port           string
	CertPath       string
	KeyPath        string
	FileServerPath string
	Authuser       string
	Authpass       string
	RequestLogPath string
	CookieLogPath  string
	RequestLogDelim string
	CookieLogDelim string

	Killswitch     bool
	Quite          bool
	Authentication bool

	// Colors and log indicators
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"

	LogSuccess = Green + "[+]" + Reset
	LogInfo    = Blue + "[i]" + Reset
	LogWarn    = Yellow + "[!]" + Reset
	LogErr     = Red + "[-]" + Reset

	// HTML
	uploadForm = `<html>
    <h3>Webpocket</h3>
    <div>
        <form enctype="multipart/form-data" method="post" action="/f">
            <input type="file" name="data"/>
            <input type="submit" value="Upload">
        </form>
    </div>
</html>`

	uploadSuccess = "<html> <h3>Success!</h3> </html>"
	uploadFail    = "<html> <h3>Error!</h3> </html>"
	illegalMethod = "<html> <h3>Illegal!</h3> </html>"
	unauthorized  = "<html> <h3>Unauthorized :(</h3> </html>"
)

func init() {

	flag.BoolVar(&HelpFlag, "h", false, "Print this message")
	flag.BoolVar(&RunCookieStealer, "c", false, "activate cookiestealer")
	flag.BoolVar(&RunRequestBin, "b", false, "activate request bin")
	flag.BoolVar(&TlsOn, "tls", false, "Activate tls")
	flag.BoolVar(&RunFileServer, "fs", false, "activate file serve")

	flag.StringVar(&CertPath, "cert", "./cert.pem", "Path to certificate")
	flag.StringVar(&KeyPath, "key", "./key.pem", "Path to key for certificate")
	flag.StringVar(&Address, "a", "0.0.0.0", "Address to listen on")
	flag.StringVar(&Port, "p", "9080", "Port\n-p 1234")
	flag.StringVar(&FileServerPath, "fspath", "./", "Path for file server")
	flag.StringVar(&RequestLogPath, "rl", "requestlog_time.txt", "Output file for request bin. -b needs to be provided")
	flag.StringVar(&CookieLogPath, "cl", "cookielog_time.txt", "Output file for cookielog. -c needs to be provided")
	flag.StringVar(&RequestLogDelim, "rld", ",", "Delimiter for request logs")
	flag.StringVar(&CookieLogDelim, "cld", ",", "Delimiter for cookie logs")

	flag.BoolVar(&Killswitch, "k", false, "killswitch, server shuts down after receiving a file")
	flag.BoolVar(&Quite, "q", false, "Suppress output")
	flag.BoolVar(&Authentication, "auth", false, "Activate basic authentication")
	flag.StringVar(&Authuser, "authuser", "webpocket", "Username for basic authentication")
	flag.StringVar(&Authpass, "authpass", "webpocket", "Password for basic authentication")

	flag.IntVar(&ParserSize, "s", 32<<20, "Max file size\n-s 200000 || -s $((2 << 20))")

	flag.Parse()

	parserFloat, parserUnit := CalcBufferSize()

	// info messages
	if !Quite {
		log.Printf("%s Listening on address: %s\n", LogInfo, Address)
		log.Printf("%s Running on port %s\n", LogInfo, Port)
		log.Printf("%s Max upload size: %.2F %s\n", LogInfo, parserFloat, parserUnit)
	}

	if Authentication {

		// check creds
		if Authuser == "" || Authpass == "" {
			if !Quite {
				log.Printf("%s Need to provide username and password when using authentication", LogErr)
			}
			os.Exit(-4)
		} else if Authuser == "webpocket" && Authpass == "webpocket" {
			if !Quite {
				log.Println("%s Using default credentials for HTTP Basic Auth!", LogWarn)
			}
		}

		// Print auth-header for requests

		if !Quite {
			data := []byte(Authuser + ":" + Authpass)
			authHeader := base64.StdEncoding.EncodeToString(data)
			fmt.Printf("Authorization header:\nAuthorization: Basic %s\n", authHeader)
		}

	}

}
