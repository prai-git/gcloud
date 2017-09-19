package main

import (
	"fmt"
	"github.com/prai-git/gcloud/cloudapp"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", cloudapp.SaveDataIntoDataStore)
	fmt.Print("Listening on port 5000")
	fmt.Print(http.ListenAndServe(":5000", nil))
}
