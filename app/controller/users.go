package controller

import (
	"html/template"
	"net/http"
	"path/filepath"

	"example.com/siteGolang/app/model"
	"github.com/julienschmidt/httprouter"
)

func GetUsers(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user, err := model.GetAllUsers()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	main := filepath.Join("app", "public", "html", "usersDynamicPage.html")
	common := filepath.Join("app", "public", "html", "common.html")

	tmpl, err := template.ParseFiles(main, common)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	err = tmpl.ExecuteTemplate(rw, "users", user)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}
