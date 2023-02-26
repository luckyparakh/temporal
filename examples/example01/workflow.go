package example01

import "go.temporal.io/sdk/workflow"

// Way to Call
func Workflow(ctx workflow.Context, a, b int) (int, error) {
	workflow.GetLogger(ctx).Info("starting example 01")
	return a + b, nil
}
