package main

import (
	"net/http"
)

func main() {
	r := getRouter()
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err.Error())
	}
}
