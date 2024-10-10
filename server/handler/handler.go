package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type myMap struct {
	m    map[string]string
	lock sync.RWMutex
}

func NewMap() {
	mapka = myMap{
		m: make(map[string]string),
	}
	fmt.Println("map is initialized!")
}

var mapka myMap

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Println("Hello, World!")

}

func HandlerMutation(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:

		var requestData map[string]string
		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		key, keyOk := requestData["key"]
		value, valueOk := requestData["value"]

		if !keyOk || !valueOk || key == "" || value == "" {
			http.Error(w, "Key and value are required and must not be empty", http.StatusBadRequest)
			return
		}
		mapka.addToMap(w, key, value)
		w.WriteHeader(http.StatusOK)

	case http.MethodDelete:
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "Key is required for deletion", http.StatusBadRequest)
			return
		}
		mapka.deleteFromMap(w, key)
		w.WriteHeader(http.StatusOK)

	case http.MethodGet:
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "Key is required", http.StatusBadRequest)
			return
		}
		value := mapka.getValueFromMap(key)
		if value == "" {
			http.Error(w, "Key is not found", http.StatusNotFound)
			return
		}
		_, err := fmt.Fprintf(w, "For key %s, value: %s\n", key, value)
		if err != nil {
			log.Printf("Encountered error: %s", err)
		}
		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, "invalid http Method", http.StatusMethodNotAllowed)
	}
}

func (m *myMap) addToMap(w http.ResponseWriter, key string, value string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.m[key] = value
	log.Printf("Added to map key %s with value %s\n", key, value)
	_, err := fmt.Fprintf(w, "Added to map key %s with value %s\n", key, value) // to client
	if err != nil {
		log.Printf("Encountered error: %s", err)
	}
}

func (m *myMap) deleteFromMap(w http.ResponseWriter, key string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	//checking if this key exists
	value, ok := m.m[key]
	if !ok {
		http.Error(w, fmt.Sprintf("Key %s not found in the map", key), http.StatusNotFound)
		return
	}

	delete(m.m, key)
	_, err := fmt.Fprintf(w, "Key %s and its value %s were deleted from the map\n", key, value) // to client
	if err != nil {
		log.Printf("Encountered error: %s", err)
		http.Error(w, fmt.Sprintf("Unexpected error: %s", err), http.StatusInternalServerError)
		return
	}

	if len(m.m) == 0 {
		log.Println("Map is empty now")
		_, err := fmt.Fprintf(w, "Map is empty now\n") // to client
		if err != nil {
			log.Printf("Encountered error: %s", err)
		}
	} else {
		log.Println("There is still something in the map")
		_, err := fmt.Fprintln(w, "There is still something in the map") // to client
		if err != nil {
			log.Printf("Encountered error: %s", err)
		}
	}
}

func (m *myMap) getValueFromMap(key string) string {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.m[key]
}
