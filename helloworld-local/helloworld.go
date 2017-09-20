package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
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

type Task struct {
	Description string
	Name        string
	Age         int
}

var tasks []Task

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	log.Printf("listening on port %s", *flagPort)
	log.Fatal(http.ListenAndServe(":"+*flagPort, mux))
}

var results []string

// GetHandler handles the index route
func GetHandler(w http.ResponseWriter, r *http.Request) {

	jsonBody, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, "Error converting results to json",
			http.StatusInternalServerError)
	}
	w.Write(jsonBody)
}

// PostHandler converts post request body to string
func handler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		postHandler(w, r)
	} else if r.Method == "GET" {
		getsHandler(w, r)
	} else {
		http.Error(w, "Invalid request method"+r.RequestURI, http.StatusMethodNotAllowed)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var t Task
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &t)

	tasks = append(tasks, t)
	//set(task);

	j, _ := json.Marshal(t)
	w.Write(j)
}

func getsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	queryValues := r.URL.Query()

	t := Task{}
	if len(queryValues.Get("Name")) > 0 {
		t.Name = queryValues.Get("Name")
	}

	if len(queryValues.Get("Age")) > 0 {
		age := queryValues.Get("Age")
		t.Age, _ = strconv.Atoi(age)
	}

	if len(queryValues.Get("Description")) > 0 {
		t.Description = queryValues.Get("Description")
	}
	if len(t.Name) > 0 {
		tasks = append(tasks, t)
		// set(t);
	}

	j, _ := json.Marshal(tasks) // get()
	w.Write(j)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func set(task Task) {
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

	var entities []Task

	q := datastore.NewQuery("helloworld")
	_, err = client.GetAll(ctx, q, &entities)

	for _, v := range entities {
		fmt.Printf("[%s] %s, %d\n", v.Description, v.Name, v.Age)
	}
}

func get() []Task {
	ctx := context.Background()

	// Set your Google Cloud Platform project ID.
	projectID := PROJECT_ID

	// Creates a client.
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	var entities []Task

	q := datastore.NewQuery("helloworld")
	_, err = client.GetAll(ctx, q, &entities)

	for _, v := range entities {
		fmt.Printf("[%s] %s, %d\n", v.Description, v.Name, v.Age)
	}

	return entities

}
