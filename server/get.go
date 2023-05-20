package server

import (
	"elix/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		return
	}
	
	authHeader := r.Header.Get("Authorization")
	authToken := os.Getenv("AUTH_TOKEN")
	if authHeader != authToken {
		http.Error(w, "Unauthorized.", http.StatusUnauthorized)
		return;
	}

	vars := mux.Vars(r)
	key := vars["key"]
	if utils.Included([]string{"favicon.ico"}, key) {
		w.Write([]byte(fmt.Sprintf("ElixDB")))
	} else {
		value := utils.FSGetViaKey("elix", key)
		jsonData, err := json.Marshal(value)
		if err != nil {
			http.Error(w, "Error parsing json to string.", http.StatusMethodNotAllowed)
			utils.Logger.Error("Error parsing json to string.")
			return
		}
		w.Write([]byte(fmt.Sprintf("%v", string(jsonData))))

	}
}
