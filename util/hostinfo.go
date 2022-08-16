package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

type FileInfo struct {
	Name string
	Size int64
}

const (
	MAIN_APP_DIR   = "xdrop"
	UPLOAD_DIR     = "uploads"
	DOWNLOADS_DIR  = "downlaods"
	DEFAULT_PERM   = 0755
	WEBSERVER_PORT = ":9080"
)

func GetClientIp(r *http.Request) string {
	// @todo use x-forwarded
	// ip := r.RemoteAddr
	// xforward := r.Header.Get("X-Forwarded-For")
	client_info := strings.Split(r.RemoteAddr, ":")
	return client_info[0]
}

func GetLocalIp() string {
	conn, error := net.Dial("udp", "8.8.8.8:80")
	if error != nil {
		fmt.Println(error)
	}

	defer conn.Close()
	server_ip := strings.Split(conn.LocalAddr().(*net.UDPAddr).String(), ":")
	return server_ip[0]
}

func GetSererPort() string {
	return WEBSERVER_PORT
}

func GetUploadDirName() string {
	return UPLOAD_DIR
}

func GetDownloadDirName() string {
	return DOWNLOADS_DIR
}

func GetMainDirName() string {
	return MAIN_APP_DIR
}

/**
* Returns the path to the home directory of the user
 */
func GetUserHomeDir() string {
	dirname, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	return dirname
}

func GetUserAppDir() string {
	return createDirIfNotExist(GetUserHomeDir(), MAIN_APP_DIR)
}

func GetUploadDirPath() string {
	return createDirIfNotExist(GetUserAppDir(), UPLOAD_DIR)
}

func GetDownloadDirPath() string {
	return createDirIfNotExist(GetUserAppDir(), DOWNLOADS_DIR)
}

func GetFileListInUpalod() []FileInfo {
	var files []FileInfo

	files, err := iOReadDir(GetUploadDirPath())
	if err != nil {
		log.Fatal("Some error occured while reading the folder")
		return files
	}

	return files
}

func GetFileListInDownload() []FileInfo {
	var files []FileInfo

	files, err := iOReadDir(GetDownloadDirPath())
	if err != nil {
		log.Fatal("Some error occured while reading the folder")
		return files
	}

	return files
}

func iOReadDir(root string) ([]FileInfo, error) {
	var files []FileInfo

	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, FileInfo{Name: file.Name(), Size: file.Size()})
	}
	return files, nil
}

/**
* Create the specified directory if it doesn't exist and return the full path to the dir.
 */
func createDirIfNotExist(base_path string, dirname string) string {

	fmt.Printf("Check if %s exists in %s\n", dirname, base_path)

	var full_path_name string = base_path + "/" + dirname

	_, err := os.Stat(full_path_name)
	if os.IsNotExist(err) {
		// create folder
		fmt.Println("folder doesn't exists, creating...")
		err := os.Mkdir(full_path_name, DEFAULT_PERM)
		if err != nil {
			log.Fatal(err)
		}
	}

	return full_path_name

}
