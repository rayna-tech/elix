package server

import (
	"elix/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		return
	}
	authHeader := r.Header.Get("Authorization")
	authToken := os.Getenv("AUTH_TOKEN")
	if authHeader != authToken {
		http.Error(w, "Unauthorized.", http.StatusUnauthorized)
		return;
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.Logger.Error("Error reading request body.")
		http.Error(w, "Error reading request body.", http.StatusBadRequest)
		return
	}

	var reqBody utils.KeyStorePair
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		utils.Logger.Error("Error parsing request body.")
		http.Error(w, "Error parsing request body.", http.StatusBadRequest)
		return
	}

	if reqBody.Key == "" {
		utils.Logger.Error("Key value is required.")
		http.Error(w, "Key value is required.", http.StatusBadRequest)
		return
	}

	if reqBody.Store == nil {
		utils.Logger.Error("Store value is required.")
		http.Error(w, "Store value is required.", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	failed := utils.FSAppendToStore("elix", reqBody.Key, reqBody.Store) 
	if failed {
		http.Error(w, "Failed to append data to store. Check the terminal for more information.", http.StatusBadRequest)
		return
	}

	w.Write([]byte(fmt.Sprintf("Store value has been written to key. Key: %s", reqBody.Key)))
}