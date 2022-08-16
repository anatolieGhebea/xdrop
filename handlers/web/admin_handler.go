package web

import (
	"fmt"
	"net/http"

	"github.com/xdrop/templates"
	"github.com/xdrop/util"
)

type ServerInfo struct {
	UploadDir   string
	DownloadDir string
	LocalIp     string
	Port        string
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {

	server_info := ServerInfo{UploadDir: util.GetUploadDirPath(), DownloadDir: util.GetDownloadDirPath(), LocalIp: util.GetLocalIp(), Port: util.GetSererPort()}

	fmt.Printf("list: %+v\n", util.GetFileListInUpalod())
	fmt.Printf("list: %+v\n", util.GetFileListInDownload())

	templates.RenderTemplate(w, "./templates/html/admin.html", server_info)
}
