package webpocket

import (
	"fmt"
	"net/http"
	_"embed"
	"flag"
	"io"
	"os"
	"log"
)

var (
	Killswitch bool
	ParserSize int

	//go:embed html/form.html
	uploadForm string

	//go:embed html/success.html
	uploadSuccess string

	//go:embed html/fail.html
	uploadFail string

	//go:embed html/illegal.html
	illegalMethod string

)

func init() {
	flag.BoolVar(&Killswitch, "k", false, "killswitch, server shuts down after receiving a file")
	flag.IntVar(&ParserSize, "s", 32<<20, "Max file size\n-s 200000 || -s $((2 << 20))")

}

func fileHandlor(writer http.ResponseWriter, req *http.Request) {

	// parse form
	err := req.ParseMultipartForm(int64(ParserSize))
	if err != nil {
		fmt.Fprint(writer, uploadFail+err.Error())
		return
	}

	// i guess the file is in memory now
	inFile, handler, err := req.FormFile("data")
	if err != nil {
		fmt.Fprint(writer, uploadFail, err.Error())
		return
	}
	defer inFile.Close()

	// create file for output
	outFile, err := os.Create(handler.Filename)
	if err != nil {
		fmt.Fprint(writer, uploadFail, err.Error())
		return
	}
	defer outFile.Close()

	// write to file, duh
	io.Copy(outFile, inFile)

	// say we good
	log.Printf("[+] File %s received!\n", handler.Filename)
	fmt.Fprint(writer, uploadSuccess)

	//check killswitch
	if Killswitch {
		os.Exit(0)
	}

	return

}


func UploadHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		fmt.Fprintf(w, uploadForm)
	case "POST":
		fileHandlor(w, req)

	default:
		fmt.Fprintf(w, illegalMethod)
	}

}

func CalcBufferSize() (float64, string) {

	if ParserSize/1e9 >= 1 {
		retVal := float64(ParserSize / 1e9)
		return retVal, "GB"
	} else if ParserSize/1e6 >= 1 {
		retVal := float64(ParserSize / 1e6)
		return retVal, "MB"
	} else if ParserSize/1000 >= 1 {
		retVal := float64(ParserSize / 1000)
		return retVal, "KB"
	} else {
		retVal := float64(ParserSize)
		return retVal, "B"
	}

}
