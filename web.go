package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handlerHelloWorld)
	http.HandleFunc("/name", handlerName)
	http.ListenAndServe(":8080", nil)
}

func handlerHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World !!!")
}

func handlerName(w http.ResponseWriter, r *http.Request) {
	names := []string{"Bryan", "Riky", "Sean", "Kenneth", "Indra"}

	namesJSON, err := json.Marshal(names)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(namesJSON)
}
