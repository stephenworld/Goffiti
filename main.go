package main

import (
	"dockerize/server/utils"
	"html/template"
	"net/http"
)

func main() {

	page, _ := template.ParseFiles("web/index.html")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			page.Execute(w, "404, Page not found")
			return
		}

		if r.Method != http.MethodPost {
			page.Execute(w, nil)
			return
		}

	})

	http.HandleFunc("/ascii-art", func(w http.ResponseWriter, r *http.Request) {
		STRING, BANNER := r.FormValue("string"), r.FormValue("banner")
		BANNER = utils.HandleBanner(BANNER)

		if STRING == "" || BANNER == "" {
			page.Execute(w, "400, Incomplete Request")
			return
		}

		art := utils.ProcessAscii(STRING, BANNER)

		if art == "" {
			page.Execute(w, "500, Server Error")
			return
		}

		page.Execute(w, art)

	})

	http.ListenAndServe(":10098", nil)
}
