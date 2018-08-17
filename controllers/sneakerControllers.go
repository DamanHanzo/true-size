package controllers

import (
	"net/http"
	"github.com/damanhanzo/true-size/models"
	"encoding/json"
	u "github.com/damanhanzo/true-size/utils"
)

var CreateSneaker = func(w http.ResponseWriter, r *http.Request) {

	sneaker := &models.Sneaker{}

	err := json.NewDecoder(r.Body).Decode(sneaker)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := sneaker.Create()
	u.Respond(w, resp)
}

var GetSneakers = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetSneakers()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}