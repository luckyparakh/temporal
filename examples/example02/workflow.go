package example02

import "go.temporal.io/sdk/workflow"

type Input struct{
	A int
	B int
}

// Way to Call
// docker exec temporal-admin-tools tctl wf run --tq workshop --et 30 --wt example01 -i '{"A":1,"B":2}'
func Workflow(ctx workflow.Context, i Input) (int, error) {
	workflow.GetLogger(ctx).Info("starting example 02")
	return i.A + i.B, nil
}
