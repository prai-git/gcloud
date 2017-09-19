package main

import (
	"fmt"
	"net/http"
	"github.com/prai-git/gcloud/helloworld"
	
)

func main() {
	http.HandleFunc("/hello", helloworld.SaveDataIntoDataStore)
	fmt.Print("Listening on port 5000")
	fmt.Print(http.ListenAndServe(":5000", nil))
}
