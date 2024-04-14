// main.go
package main

import (
	"fmt"
	"golang-cicd/config"
	"net/http"

	"github.com/kataras/golog"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	if !config.Parse("conf/conf.yaml") {
		golog.Error("Invalid Config provided")
		return
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(config.Props.Listen, nil)
}
