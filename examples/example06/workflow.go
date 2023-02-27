package example06

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type WorkflowInput struct {
	A int
	B int
}

type WorkflowOutput struct {
	Result int
}

// Code in workflow should be deterministic
func Workflow(ctx workflow.Context, i WorkflowInput) (WorkflowOutput, error) {
	workflow.GetLogger(ctx).Info("starting example 06")
	rp := &temporal.RetryPolicy{
		InitialInterval: 1,
		BackoffCoefficient: 2,
		MaximumInterval: 100,
		MaximumAttempts: 5,
	}
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		TaskQueue:              "workshop",
		ScheduleToCloseTimeout: 3*time.Second + 3*time.Second,
		ScheduleToStartTimeout: 3 * time.Second,
		StartToCloseTimeout:    3 * time.Second,
		HeartbeatTimeout:       0 * time.Second,
		WaitForCancellation:    false,
		ActivityID:             "",
		RetryPolicy:            rp,
	})
	var output Output
	ii := Input{A: i.A, B: i.B}
	err := workflow.ExecuteActivity(ctx, Activity06, ii).Get(ctx, &output)
	if err != nil {
		return WorkflowOutput{}, err
	}
	return WorkflowOutput{Result: output.C}, nil
}
