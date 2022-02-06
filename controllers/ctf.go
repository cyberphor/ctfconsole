package controllers

import (
	"net/http"
)

func CreateCtfPage(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "create-ctf.gohtml", nil)
}
