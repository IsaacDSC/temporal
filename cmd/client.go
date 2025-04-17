package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.temporal.io/sdk/temporal"
	"log"
	"time"
	"workflows/internal/service"

	"go.temporal.io/sdk/client"
)

func main() {
	// The client is a heavyweight object that should be created once per process.
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID: uuid.New().String(),
		//ID:        "6b44d13c-251c-4090-bc60-79683cd1ad7e",
		TaskQueue: "transaction:v2",
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    time.Second * 100,
			MaximumAttempts:    5,
		},
	}

	input := service.Input{Name: "input.Temporal"}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, service.Workflow, &input)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	fmt.Println()
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
	fmt.Println()

	// Synchronously wait for the workflow completion.
	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow result:", result)
}
