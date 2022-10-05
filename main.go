package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type ServiceResult struct {
	FormattedTime string
}

func currentTime() string {
	return time.Now().Format(time.RFC822Z)
}

func handler(w http.ResponseWriter, r *http.Request) {
	t := currentTime()
	sr := ServiceResult{t}
	js, err := json.Marshal(sr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
