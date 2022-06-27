package handlers

import (
	"net/http"
	"encoding/json"
)

func TestFunc(w http.ResponseWriter, r *http.Request){
		wel := map[string]string{
			"Greetings":"Hello World!",
			"Status":"ok",
		}
		w.Header().Set("Content-Type","application/json")
		json.NewEncoder(w).Encode(wel)
}