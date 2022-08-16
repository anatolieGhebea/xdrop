package handlers

import (
	"github.com/gorilla/mux"
	"github.com/xdrop/handlers/api"
	"github.com/xdrop/handlers/web"
	"github.com/xdrop/middlewares"
)

func GetRouter() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/", web.HomeHandler)
	r.HandleFunc("/403.html", web.HomeHandler)
	r.HandleFunc("/admin", middlewares.AllowOnlyLocalhost(web.AdminHandler))
	r.HandleFunc("/stop_server", middlewares.AllowOnlyLocalhost(web.StopServerHandler))

	rest_api := r.PathPrefix("/api/").Subrouter()

	rest_api.HandleFunc("/upload_file", api.UplaodFileHandler).Methods("POST")
	rest_api.HandleFunc("/download_file/{filename}", api.DownloadFileHandler).Methods("GET")
	rest_api.HandleFunc("/get_downloadable_file_list", api.DownloadableListHandler).Methods("GET")
	rest_api.HandleFunc("/get_uploaded_file_list", api.UploadedListHandler).Methods("GET")

	return r
}
