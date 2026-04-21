package main

import (
	"fmt"
	"html/template"
	"net/http"
	"stylize/server/utils"
)

type DATA struct {
	ART   string
	INPUT string
	FONT  string
}

func main() {
	page, _ := template.ParseFiles("web/index.html")
	errorPage, _ := template.ParseFiles("web/404.html")

	fmt.Println("Server live at http://localhost:8080/")

	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/web/", http.StripPrefix("/web/", fs))

	STRING, BANNER := "", ""

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			art := utils.GenerateAsciiArt("Error 404.\nPage not found", "server/banners/standard.txt")
			errorPage.Execute(w, art)
			return
		}

		if r.Method != http.MethodPost {
			page.Execute(w, nil)
			return
		}

	})

	http.HandleFunc("/ascii-art", func(w http.ResponseWriter, r *http.Request) {
		STRING, BANNER = r.FormValue("input"), r.FormValue("font")
		BANNER = utils.HandleFont(BANNER)

		if BANNER == "" || STRING == "" {
			art := utils.GenerateAsciiArt("Error 400.\nBad Request", "server/banners/standard.txt")
			page.Execute(w, DATA{ART: art, INPUT: STRING, FONT: BANNER})
			return
		}

		art := utils.GenerateAsciiArt(STRING, BANNER)
		if art == "" {
			art := utils.GenerateAsciiArt("Error 500.\nServer Error", "server/banners/standard.txt")
			page.Execute(w, DATA{ART: art, INPUT: STRING, FONT: BANNER})
			return
		}
		page.Execute(w, DATA{ART: art, INPUT: STRING, FONT: BANNER})
	})

	http.ListenAndServe(":8080", nil)
}
