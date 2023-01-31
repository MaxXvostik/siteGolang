package controller

import (
	"encoding/json"
	"fmt"
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

	main := filepath.Join("public", "html", "usersDynamicPage.html")
	common := filepath.Join("public", "html", "common.html")

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

func AddUser(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := r.FormValue("name")

	surname := r.FormValue("surname")

	fmt.Println(name)

	if name == "" || surname == "" {
		http.Error(rw, "Имя и фамилия не могут быть пустыми", 400)
		return
	}
	user := model.NewUser(name, surname)
	err := user.Add()
	if err != nil {
		http.Error(rw, err.Error(), 400)
	}
	err = json.NewEncoder(rw).Encode("Пользователь успешно добавлен!")
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}
