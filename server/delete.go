package server

import (
	// "elix/utils"
	// "fmt"
	"elix/utils"
	"fmt"
	"net/http"
	"os"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	authHeader := r.Header.Get("Authorization")
	authToken := os.Getenv("AUTH_TOKEN")
	if authHeader != authToken {
		http.Error(w, "Unauthorized.", http.StatusUnauthorized)
		return;
	}
	if key == "" {
		utils.Logger.Error("Key value is required.")
		http.Error(w, "Key value is required.", http.StatusBadRequest)
		return
	}
	if utils.FSStoreContainsKey("elix", key) {
		utils.Logger.Error(fmt.Sprintf("Store doesnt contain key: %s", key))
		http.Error(w, fmt.Sprintf("Store doesnt contain key: %s. Check the terminal for more information.", key), http.StatusBadRequest)
		return
	}
	utils.FSDeleteKeyFromStore("elix", key)
	w.Write([]byte(fmt.Sprintf("Deleted key & store from db. Key: %s", key)))
	return
}

func Clear(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	authToken := os.Getenv("AUTH_TOKEN")
	if authHeader != authToken {
		http.Error(w, "Unauthorized.", http.StatusUnauthorized)
		return;
	}
	utils.FSClearStore("elix")
	w.Write([]byte("Store file has been cleared/reset."))
	return
}
