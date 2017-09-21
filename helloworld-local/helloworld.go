package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"cloud.google.com/go/datastore"

	"golang.org/x/net/context"
)

var (
	// flagPort is the open port the application listens on
	flagPort = flag.String("port", "8000", "Port to listen on")
)

const PROJECT_ID string = "helloworld-180317"
const KIND string = "helloworld"

type QueryData struct {
	Key   string
	Value []string
}

var tasks []interface{}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/save", saveHandler)
	mux.HandleFunc("/retrive", retriveHandler)

	log.Printf("listening on port %s", *flagPort)
	log.Fatal(http.ListenAndServe(":"+*flagPort, mux))
}

func saveHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		w.Header().Set("Content-Type", "application/json")
		queryValues := r.URL.Query()
		fmt.Printf("%T", queryValues)

		for key, value := range queryValues {
			q := QueryData{key, value}
			fmt.Println("Key:", key, "Value:", value)
			tasks = append(tasks, q)
		}

		j, _ := json.Marshal("Saved Data Successfull") // get()
		w.Write(j)
	} else {
		http.Error(w, "Invalid request method"+r.RequestURI, http.StatusMethodNotAllowed)
	}
}

func retriveHandler(w http.ResponseWriter, r *http.Request) {

	j, _ := json.Marshal(tasks) // get()
	w.Write(j)
}

func set(task interface{}) {
	ctx := context.Background()

	// Set your Google Cloud Platform project ID.
	projectID := PROJECT_ID

	// Creates a client.
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the kind for the new entity.
	kind := KIND
	// Sets the name/ID for the new entity.
	name := strconv.Itoa(123 + rand.Intn(100000))
	// Creates a Key instance.
	taskKey := datastore.NameKey(kind, name, nil)

	// Saves the new entity.
	if _, err := client.Put(ctx, taskKey, &task); err != nil {
		log.Fatalf("Failed to save task: %v", err)
	}

}

func get() []interface{} {
	ctx := context.Background()

	// Set your Google Cloud Platform project ID.
	projectID := PROJECT_ID

	// Creates a client.
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	var entities []interface{}

	q := datastore.NewQuery("helloworld")
	_, err = client.GetAll(ctx, q, &entities)

	return entities

}
