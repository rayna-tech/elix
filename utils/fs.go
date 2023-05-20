package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func FSAssureStoreExist(storeName string) {
	sN := fmt.Sprintf("%s.json", storeName)
	_, err := os.Stat(sN)
	if err != nil {
		data := make(map[string]interface{})
		jsonData, _ := json.Marshal(data)

		err := ioutil.WriteFile(sN, jsonData, 0644)
		if err != nil {
			Logger.Error("Failed to create inital store file.")
			return
		}
	} else if err != nil {
		Logger.Error("Error checking to see if store already exists.")
		return
	}
}

func FSAppendToStore(storeName string, key string, value interface{}) bool {
	sN := fmt.Sprintf("%s.json", storeName)
	data, err := ioutil.ReadFile(sN)
	if err != nil {
		Logger.Error("Failed to read store.")
		return true
	}

	var stores map[string]interface{}
	err = json.Unmarshal(data, &stores)
	if err != nil {
		Logger.Error("Error unmarshaling store.")
		return true
	}
	if _, ok := stores[key]; ok {
		Logger.Error(fmt.Sprintf("Store already contains key. Key: %s", key))
		return true
	}
	stores[key] = value
	jsonData, err := json.MarshalIndent(stores, "", " ")
	if err != nil {
		Logger.Error("Error remarshaling store.")
		return true
	}

	err = ioutil.WriteFile(sN, jsonData, 0644)
	if err != nil {
		Logger.Error("Failed to append to store.")
		return true
	}

	Logger.Info(fmt.Sprintf("Appended new value to store. Key: %s", key))
	return false
}

func FSGetViaKey(storeName string, key string) interface{} {
	sN := fmt.Sprintf("%s.json", storeName)
	data, err := ioutil.ReadFile(sN)
	if err != nil {
		Logger.Error("Failed to read store.")
		return nil
	}

	var stores map[string]interface{}
	err = json.Unmarshal(data, &stores)
	if err != nil {
		Logger.Error("Error unmarshaling store.")
		return nil
	}

	if _, ok := stores[key]; ok == false{
		Logger.Error(fmt.Sprintf("Store doesnt contain key. Key: %s", key))
		return nil
	}

	return stores[key]
}

func FSStoreContainsKey(storeName string, key string) bool {
	sN := fmt.Sprintf("%s.json", storeName)
	data, err := ioutil.ReadFile(sN)
	if err != nil {
		Logger.Error("Failed to read store.")
		return true
	}

	var stores map[string]interface{}
	err = json.Unmarshal(data, &stores)
	if err != nil {
		Logger.Error("Error unmarshaling store.")
		return true
	}

	if _, ok := stores[key]; ok == false{
		Logger.Error(fmt.Sprintf("Store doesnt contain key. Key: %s", key))
		return true
	}

	return false
}

func FSDeleteKeyFromStore(storeName string, key string) {
	sN := fmt.Sprintf("%s.json", storeName)
	data, err := ioutil.ReadFile(sN)
	if err != nil {
		Logger.Error("Failed to read store.")
		return
	}

	var stores map[string]interface{}
	err = json.Unmarshal(data, &stores)
	if err != nil {
		Logger.Error("Error unmarshaling store.")
		return 
	}

	delete(stores, key)
	jsonData, err := json.MarshalIndent(stores, "", " ")
	if err != nil {
		Logger.Error("Error remarshaling store.")
		return
	}

	err = ioutil.WriteFile(sN, jsonData, 0644)
	if err != nil {
		Logger.Error("Failed to update store after deletion.")
		return
	}

	Logger.Info(fmt.Sprintf("Deleted store value. Key: %s", key))
	return 
}

func FSClearStore(storeName string) {
	sN := fmt.Sprintf("%s.json", storeName)
	file, err := os.OpenFile(sN, os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		Logger.Error("Failed to open store file.")
		return
	}	
	defer file.Close()
	_, err = file.WriteString("{}")
	if err != nil {
		Logger.Error("Failed to clear store file.")
		return
	}
	Logger.Info("Store file has been cleared/reset.")
}

func FSUpdateValue(storeName string, key string, value interface{}) bool {
	sN := fmt.Sprintf("%s.json", storeName)
	data, err := ioutil.ReadFile(sN)
	if err != nil {
		Logger.Error("Failed to read store.")
		return true
	}

	var stores map[string]interface{}
	err = json.Unmarshal(data, &stores)
	if err != nil {
		Logger.Error("Error unmarshaling store.")
		return true
	}
	if _, ok := stores[key]; ok == false {
		Logger.Error(fmt.Sprintf("Store doesnt contain key. Key: %s", key))
		return true
	}
	stores[key] = value
	jsonData, err := json.MarshalIndent(stores, "", " ")
	if err != nil {
		Logger.Error("Error remarshaling store.")
		return true
	}

	err = ioutil.WriteFile(sN, jsonData, 0644)
	if err != nil {
		Logger.Error("Failed to update store.")
		return true
	}

	Logger.Info(fmt.Sprintf("Updated previous key value. Key: %s", key))
	return false
}