package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/xdrop/handlers"
	"github.com/xdrop/middlewares"
	"github.com/xdrop/util"
)

func open_browser(url string) {
	var err error

	fmt.Println("running on :" + runtime.GOOS)

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var dir string

	flag.StringVar(&dir, "dir", "./static", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	fmt.Printf("Starting server\n")

	// load the Router
	r := handlers.GetRouter()

	// serve static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	http.Handle("/", middlewares.PanicRecoveryHandler(r))

	// open the default browser at the /admin page
	open_browser("http://localhost" + util.GetSererPort() + "/admin")

	log.Fatal(http.ListenAndServe(util.GetSererPort(), nil))
}
