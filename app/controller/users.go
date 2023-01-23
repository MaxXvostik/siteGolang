package controller

import (
	"encoding/json"
	"net/http"

	"example.com/siteGolang/app/model"
	"github.com/julienschmidt/httprouter"
)

func GetUsers(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user, err := model.GetAllUsers()

	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	err = json.NewEncoder(rw).Encode(user)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

}
