package serve

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/cyberphor/ctfconsole/create"
	"github.com/cyberphor/ctfconsole/read"
)

func submitRegistration(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/login.html", http.StatusSeeOther)
		return
	} else {
		create.Player(r.FormValue("username"), r.FormValue("password"), r.FormValue("team"))
		http.Redirect(w, r, "/login.html", http.StatusSeeOther)
		return
	}
}

func submitLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else {
		username := r.FormValue("username")
		password := r.FormValue("password")
		fmt.Println(username, password)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

func showScoreboard(w http.ResponseWriter, r *http.Request) {
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

	http.HandleFunc("/scoreboard.html", showScoreboard)
	http.HandleFunc("/submitRegistration", submitRegistration)
	http.HandleFunc("/submitLogin", submitLogin)
	http.ListenAndServe(":8000", nil)
}
