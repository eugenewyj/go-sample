package handlers

import (
	"net/http"
	"encoding/json"
)

func Routes()  {
	http.HandleFunc("/sendjson", SendJSON)
}

func SendJSON(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name string
		Email string
	}{
		Name: "Bill",
		Email: "Bill@email.com",
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(&u)
}
