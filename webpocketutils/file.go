package webpocket

import (
	"fmt"
	"net/http"
	"io"
	"os"
	"log"
	"strings"
)

type FileHandler struct {
	OutDir string
}

func (fh *FileHandler) saveFileFromRequest(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(int64(ParserSize))
	if err != nil {
		fmt.Fprint(w, uploadFail+err.Error())
		return
	}

	inFile, handler, err := r.FormFile("data")
	if err != nil {
		fmt.Fprint(w, uploadFail, err.Error())
		return
	}
	defer inFile.Close()

	var writePath string
	if strings.HasSuffix(fh.OutDir, "/") {
		writePath = fh.OutDir + handler.Filename
	} else {
		writePath = fh.OutDir + "/" + handler.Filename
	}

	outFile, err := os.Create(writePath)
	if err != nil {
		fmt.Fprint(w, uploadFail, err.Error())
		return
	}
	defer outFile.Close()

	io.Copy(outFile, inFile)

	if !Quite {
		log.Printf("%s File %s received!\n", LogSuccess, handler.Filename)
	}
	fmt.Fprint(w, uploadSuccess)

	if Killswitch {
		os.Exit(0)
	}

	return

}


func (fh *FileHandler) UploadHandler(w http.ResponseWriter, r *http.Request) {
	if Authentication {
		if !checkAuth(r) {
			fmt.Fprintf(w, unauthorized)
			return
		}
	}

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, uploadForm)
	case "POST":
		fh.saveFileFromRequest(w, r)

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
