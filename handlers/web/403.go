package web

import (
	"net/http"

	"github.com/xdrop/templates"
)

func PermissionDeniedHandler(w http.ResponseWriter, r *http.Request) {

	templates.RenderTemplate(w, "./templates/html/403.html", nil)
}
