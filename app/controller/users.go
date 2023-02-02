package controller

import (
	"encoding/json"
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

func DeleteUser(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId := p.ByName("userId")

	user, err := model.GetUserById(userId)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	err = user.Delete()
	if err != nil {
		http.Error(rw, err.Error(), 400)
	}
	err = json.NewEncoder(rw).Encode("Пользователь был удален")
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}

func UpdateUser(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//получаем значение из параметра userId, переданного в строке запроса
	userId := p.ByName("userId")
	//получаем значения из параметров name и surname, переданных в форме запроса
	name := r.FormValue("name")
	surname := r.FormValue("surname")

	//получаем пользователя из БД по его id
	user, err := model.GetUserById(userId)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	//заменяем старые значения на новые
	user.Name = name
	user.Surname = surname

	//обновляем данные в таблице
	err = user.Update()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	//возвращаем текстовое подтверждение об успешном выполнении операции
	err = json.NewEncoder(rw).Encode("Пользователь был успешно изменен")
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}
