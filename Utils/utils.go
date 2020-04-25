package Utils

import (
	"books-rest-api/Models"
	"encoding/json"
	"net/http"
)

func SendError(w http.ResponseWriter, status int, err Models.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

func SendSucces(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}