package webpocket

import (
	"fmt"
	"net/http"
	"io"
	"os"
	"log"
)

func fileHandlor(w http.ResponseWriter, r *http.Request) {

	// parse form
	err := r.ParseMultipartForm(int64(ParserSize))
	if err != nil {
		fmt.Fprint(w, uploadFail+err.Error())
		return
	}

	// i guess the file is in memory now
	inFile, handler, err := r.FormFile("data")
	if err != nil {
		fmt.Fprint(w, uploadFail, err.Error())
		return
	}
	defer inFile.Close()

	// create file for output
	outFile, err := os.Create(handler.Filename)
	if err != nil {
		fmt.Fprint(w, uploadFail, err.Error())
		return
	}
	defer outFile.Close()

	// write to file, duh
	io.Copy(outFile, inFile)

	// say we good
	if !Quite {
		log.Printf("[+] File %s received!\n", handler.Filename)
	}
	fmt.Fprint(w, uploadSuccess)

	//check killswitch
	if Killswitch {
		os.Exit(0)
	}

	return

}


func UploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, uploadForm)
	case "POST":
		fileHandlor(w, r)

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
