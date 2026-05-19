package main

import (
	"dockerize/server/utils"
	"fmt"
	"html/template"
	"net/http"
)

type data struct {
	ART   string
	FONT  string
	INPUT string
}

func main() {
	page, err := template.ParseFiles("web/index.html")
	fmt.Println("Server live at http://localhost:8080/")

	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/web/", http.StripPrefix("/web/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Handling Server Error.
		if err != nil {
			art := utils.GenerateAsciiArt("Error 500.\nServer Error", "server/banners/standard.txt")
			page.Execute(w, data{ART: art})
			return
		}

		// Handling 404 Error
		if r.URL.Path != "/" {
			art := utils.GenerateAsciiArt("Error 404.\nPage not found", "server/banners/standard.txt")
			page.Execute(w, data{ART: art})
			return
		}

		// Rendering The Page without any update
		if r.Method != http.MethodPost {
			page.Execute(w, nil)
			return
		}

	})

	http.HandleFunc("/ascii-art", func(w http.ResponseWriter, r *http.Request) {
		// Handling Server Error
		if err != nil {
			art := utils.GenerateAsciiArt("Error 500.\nServer Error", "server/banners/standard.txt")
			page.Execute(w, data{ART: art})
			return
		}

		STRING, BANNER := r.FormValue("input"), r.FormValue("font")
		BANNER = utils.HandleFont(BANNER)
		if r.Method != http.MethodPost {
			art := utils.GenerateAsciiArt("Error 405.\nMethod Not Allowed", "server/banners/standard.txt")
			page.Execute(w, data{ART: art})
			return
		}

		if BANNER == "" || STRING == "" {
			art := utils.GenerateAsciiArt("Error 400.\nBad Request", "server/banners/standard.txt")
			page.Execute(w, data{ART: art})
			return
		}

		art := utils.GenerateAsciiArt(STRING, BANNER)
		if art == "" {
			art := utils.GenerateAsciiArt("Error 500.\nServer Error", "server/banners/standard.txt")
			page.Execute(w, data{ART: art})
			return
		}
		page.Execute(w, data{ART: art, INPUT: STRING, FONT: BANNER})
	})

	http.ListenAndServe(":8080", nil)
}
