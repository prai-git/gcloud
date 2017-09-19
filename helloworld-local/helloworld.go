// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Sample helloworld is a basic App Engine flexible app.
package main

import (
	"fmt"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	
	"cloud.google.com/go/datastore"

	"golang.org/x/net/context"
	
)


func main() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/_ah/health", healthCheckHandler)
	log.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	
	//mainFun()
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello world!")
	mainFun2(w, r)
	//handleVisit(w, r)
	fmt.Fprint(w, "Hello world! success")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

type Task struct {
	Description string
	Name string
	Age int
}

func mainFun2(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Set your Google Cloud Platform project ID.
	projectID := "helloworld-180317"

	// Creates a client.
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the kind for the new entity.
	kind := "helloworld"
	// Sets the name/ID for the new entity.
	 name := strconv.Itoa(123 + rand.Intn(100000))
	// Creates a Key instance.
	taskKey := datastore.NameKey(kind, name, nil)

	// Creates a Task instance.
	task := Task{
		Description: "datastore app",
		Name : "pankaj",
		Age: 25,
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
	
	// fmt.Fprint(w, "Hello, world!")
	
}