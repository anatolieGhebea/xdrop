package middlewares

import (
	"fmt"
	"net/http"

	"github.com/xdrop/util"
)

func AllowOnlyLocalhost(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		client_ip := util.GetClientIp(r)

		fmt.Printf("request form %s\n", client_ip)
		if client_ip != "127.0.0.1" && client_ip != "localhost" {
			http.Redirect(w, r, "/403", 403)
			return
		}

		next.ServeHTTP(w, r)
	})
}
