package web

import (
	"net/http"
	"os"
)

func StopServerHandler(w http.ResponseWriter, r *http.Request) {

	defer os.Exit(0)

	w.Write([]byte("Server stoped!"))
}
