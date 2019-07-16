package main

import (
	"fmt"
	"net/http"
)

func greeting(text string) string {
	return fmt.Sprintf("<b>%s</b>", text)
}

func main() {
 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, greeting("Code.education Rocks!"))
 	})
 	http.ListenAndServe(":8000", nil)
}