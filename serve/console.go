package serve

import (
	"html/template"
	"net/http"

	"github.com/cyberphor/ctfconsole/read"
)

func playersPage(w http.ResponseWriter, r *http.Request) {
	page, err := template.ParseFiles("./templates/players.html")
	if err != nil {
		panic(err)
	}
	players := read.Players()
	page.Execute(w, players)
}

func Console() {
	filePath := http.Dir("./templates/")
	fileServer := http.FileServer(filePath)
	http.Handle("/", fileServer)
	playersPageHandler := http.HandlerFunc(playersPage)
	http.Handle("/players.html", playersPageHandler)
	http.ListenAndServe(":8000", nil)
}
