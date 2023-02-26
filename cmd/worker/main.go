package main

import (
	"log"
	"temporal/examples/example01"
	"temporal/examples/example02"
	"temporal/examples/example03"
	"temporal/examples/example04"
	"temporal/examples/example05"
	"temporal/examples/example06"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
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
	w := worker.New(c, taskQueue, worker.Options{})
	register(w)
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start worker", err)
	}

}

func register(w worker.Worker) {
	// w.RegisterWorkflow(example01.Workflow)
	w.RegisterWorkflowWithOptions(example01.Workflow, workflow.RegisterOptions{Name: "example01"})
	w.RegisterWorkflowWithOptions(example02.Workflow, workflow.RegisterOptions{Name: "example02"})
	w.RegisterWorkflowWithOptions(example03.Workflow, workflow.RegisterOptions{Name: "example03"})
	w.RegisterWorkflowWithOptions(example04.Workflow, workflow.RegisterOptions{Name: "example04"})
	w.RegisterWorkflowWithOptions(example05.Workflow, workflow.RegisterOptions{Name: "example05"})
	w.RegisterWorkflowWithOptions(example06.Workflow, workflow.RegisterOptions{Name: "example06"})
	w.RegisterActivity(example06.Activity06)
}
