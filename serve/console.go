package serve

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/cyberphor/ctfconsole/read"
)

func registerPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else {
		fmt.Println("username:", r.FormValue("username"))
		fmt.Println("password:", r.FormValue("password"))
	}
}

func scoreboardPage(w http.ResponseWriter, r *http.Request) {
	page, err := template.ParseFiles("./templates/scoreboard.html")
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

	scoreboardPageHandler := http.HandlerFunc(scoreboardPage)
	http.Handle("/scoreboard.html", scoreboardPageHandler)

	registerPageHandler := http.HandlerFunc(registerPage)
	http.Handle("/register.html", registerPageHandler)
	http.ListenAndServe(":8000", nil)
}
