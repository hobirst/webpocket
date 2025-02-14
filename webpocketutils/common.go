package webpocket

import (
	"log"
	"flag"
	"os"
	_ "embed"
)

var (
	Killswitch     bool
	Quite          bool
	authentication bool
	authuser       string
	authpass       string
	ParserSize     int

	requestlog string
	cookielog  string

	//go:embed html/form.html
	uploadForm string

	//go:embed html/success.html
	uploadSuccess string

	//go:embed html/fail.html
	uploadFail string

	//go:embed html/illegal.html
	illegalMethod string

	//go:embed html/unauthorized.html
	unauthorized string
)

func init() {
	flag.StringVar(&requestlog, "rl", "requestlog.txt", "Output file for request bin. -b needs to be provided")
	flag.StringVar(&cookielog, "cl", "cookielog.txt", "Output file for cookielog. -c needs to be provided")

	flag.BoolVar(&Killswitch, "k", false, "killswitch, server shuts down after receiving a file")
	flag.BoolVar(&Quite, "q", false, "Suppress output")
	flag.BoolVar(&authentication, "auth", false, "Activate basic authentication")
	flag.StringVar(&authuser, "authuser", "", "Username for basic authentication")
	flag.StringVar(&authpass, "authpass", "", "Password for basic authentication")

	flag.IntVar(&ParserSize, "s", 32<<20, "Max file size\n-s 200000 || -s $((2 << 20))")

	if authentication {
		if authuser == "" || authpass == "" {
			if !Quite {
				log.Printf("[-] Need to provide username and password when using authentication")
			}	
			os.Exit(-4)
		}
	}
}
