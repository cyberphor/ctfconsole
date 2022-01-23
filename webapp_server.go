package main

import (
	"html/template"
	"net/http"
)

func PlayersPageHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := ReadPlayersTable()
	if err != nil {
		panic(err)
	}
	type PlayerData struct {
		Username string
		Password string
	}
	Players := []PlayerData{}
	for rows.Next() {
		Player := PlayerData{}
		rows.Scan(&Player.Username, &Player.Password)
		Players = append(Players, Player)
	}
	temp, err := template.ParseFiles("players.html")
	temp.Execute(w, Players)
}

func ServeWebApp() {
	http.Handle("/", http.FileServer(http.Dir("./webapp")))
	http.Handle("/players", PlayersPageHandler)
	http.ListenAndServe(":8000", nil)
}
