package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"go.temporal.io/sdk/client"
)

const (
	address   = "127.0.0.1:7233"
	namespace = "default"
	taskQueue = "workshop"
)

func main() {
	c, err := client.Dial(client.Options{
		HostPort:  address,
		Namespace: namespace,
	})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "my-first-workflow",
		TaskQueue: taskQueue,
	}
	i1, _ := strconv.Atoi(os.Args[1])
	i2, _ := strconv.Atoi(os.Args[2])
	wr, err := c.ExecuteWorkflow(context.Background(), options, "example01", i1, i2)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", wr.GetID(), "RunID", wr.GetRunID())
	var r int
	err = wr.Get(context.Background(), &r)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow result:", r)
}
