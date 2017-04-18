package main

import (
	"net/http"
	"encoding/json"
)

type Course struct {
Name    string
Keywords []string
}
func main() {
http.HandleFunc("/", foo)
http.ListenAndServe(":3000", nil)
}
func foo(w http.ResponseWriter, r *http.Request) {
course := course{"IS-105", []string{"operating system", "data communication"}}
js, err := json.Marshal(profile)
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}
w.Header().Set("Content-Type", "application/json")
w.Write(js)
}