package example03

import "go.temporal.io/sdk/workflow"

type Input struct {
	A int
	B int
}
type Output struct {
	Result int
}

func Workflow(ctx workflow.Context, i Input) (Output, error) {
	workflow.GetLogger(ctx).Info("starting example 03")
	return Output{i.A + i.B}, nil
}
