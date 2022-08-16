package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/xdrop/util"
)

type respomseObject struct {
	Status  bool
	Message string
	Data    []util.FileInfo
}

func UplaodFileHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("selectedFile")

	if err != nil {
		fmt.Println("error in uploading", err)
		return
	}

	defer file.Close()

	fmt.Printf("file %+v\n", handler.Filename)
	fmt.Printf("size %+v\n", handler.Size)
	fmt.Printf("MIME Header %+v\n", handler.Header)

	destFile, err := os.Create(util.GetUploadDirPath() + "/" + handler.Filename)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer destFile.Close()

	if _, err := io.Copy(destFile, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	resp := respomseObject{Status: true, Message: "uplaoded"}

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResp)
}

func DownloadFileHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	fileName := params["filename"]
	fmt.Printf(fileName)
	// check if file is valid
	if len(fileName) == 0 {

		w.Header().Set("Content-Type", "application/json")

		resp := respomseObject{Status: false, Message: "file not found"}

		jsonResp, err := json.Marshal(resp)

		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		w.Write(jsonResp)
		return
	}

	http.ServeFile(w, r, util.GetDownloadDirPath()+"/"+fileName)
}

func DownloadableListHandler(w http.ResponseWriter, r *http.Request) {

	files := util.GetFileListInDownload()
	resp := respomseObject{Status: true, Message: "List of the files in the uplaod folder", Data: files}

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResp)
}

func UploadedListHandler(w http.ResponseWriter, r *http.Request) {

	files := util.GetFileListInUpalod()
	resp := respomseObject{Status: true, Message: "List of the files in the uplaod folder", Data: files}

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResp)
}
