package server

import (
	"elix/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body.", http.StatusBadRequest)
		return
	}

	var reqBody utils.KeyStorePair
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		http.Error(w, "Error parsing request body.", http.StatusBadRequest)
		return
	}

	if reqBody.Key == "" {
		http.Error(w, "Key value is required.", http.StatusBadRequest)
		return
	}

	if reqBody.Store == nil {
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