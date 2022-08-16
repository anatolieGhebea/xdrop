package web

import (
	"net/http"

	"github.com/xdrop/templates"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	templates.RenderTemplate(w, "./templates/html/index.html", nil)
}
