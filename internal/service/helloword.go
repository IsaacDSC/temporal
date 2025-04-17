package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

type Input struct {
	Name string
}

// Workflow is a Hello World workflow definition.
func Workflow(ctx workflow.Context, input *Input) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("HelloWorld workflow started", "name", input.Name)

	fmt.Println("@@@@@", input.Name)

	return "", errors.New("generate error")
	var result string
	err := workflow.ExecuteActivity(ctx, Activity, input).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}

	logger.Info("HelloWorld workflow completed.", "result", result)

	return result, nil
}

func Activity(ctx context.Context, input Input) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity", "name", input.Name)
	return "Hello " + input.Name + "!", nil
}
