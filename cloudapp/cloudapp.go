package cloudapp

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"cloud.google.com/go/datastore"

	"golang.org/x/net/context"
)

const PROJECT_ID string = "helloworld-180317"
const KIND string = "helloworld"

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
	SaveDataIntoDataStore(w, r)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

type Task struct {
	Description string
	Name        string
	Age         int
}

func SaveDataIntoDataStore(w http.ResponseWriter, r *http.Request) {
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
	// Set some random number
	name := strconv.Itoa(123 + rand.Intn(100000))
	// Creates a Key instance.
	taskKey := datastore.NameKey(kind, name, nil)

	// Creates a Task instance.
	task := Task{
		Description: "datastore app",
		Name:        "pankaj",
		Age:         25,
	}

	// Saves the new entity.
	if _, err := client.Put(ctx, taskKey, &task); err != nil {
		log.Fatalf("Failed to save task: %v", err)
	}

	var entities []Task

	q := datastore.NewQuery("helloworld")
	_, err = client.GetAll(ctx, q, &entities)

	for _, v := range entities {
		fmt.Fprintf(w, "[%s] %s, %d\n", v.Description, v.Name, v.Age)
	}
}
