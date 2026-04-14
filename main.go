package main

import (
	"fmt"
	"html/template"
	"net/http"
	"stylize/server/utils"
)

func main() {
	page, _ := template.ParseFiles("web/index.html")
	fmt.Println("Server live at http://localhost:8080/")

	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/web/", http.StripPrefix("/web/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			data := utils.GenerateAsciiArt("Error 404.\nPage not found", "server/banners/standard.txt")
			page.Execute(w, data)
			return
		}

		if r.Method != http.MethodPost {
			page.Execute(w, nil)
			return
		}

		STRING, BANNER := r.FormValue("input"), r.FormValue("font")
		BANNER = utils.HandleFont(BANNER)

		if BANNER == "" || STRING == "" {
			data := utils.GenerateAsciiArt("Error 400.\nBad Request", "server/banners/standard.txt")
			page.Execute(w, data)
			return
		}

		art := utils.GenerateAsciiArt(STRING, BANNER)
		if art == "" {
			data := utils.GenerateAsciiArt("Error 500.\nServer Error", "server/banners/standard.txt")
			page.Execute(w, data)
			return
		}
		page.Execute(w, art)
	})

	http.ListenAndServe(":8080", nil)
}
