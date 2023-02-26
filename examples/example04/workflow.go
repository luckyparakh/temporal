package example04

import (
	"math/rand"
	"time"

	"go.temporal.io/sdk/workflow"
)

type Input struct {
	Num int
}
type Output struct {
	Result int
}

var max = 10

// Code in workflow should be deterministic
func Workflow(ctx workflow.Context, in Input) (Output, error) {
	workflow.GetLogger(ctx).Info("starting example 04")
	r := 1
	number := in.Num
	if number < 1 {
		workflow.GetLogger(ctx).Info("invalid number")

		// Side Effect
		encodedNUmber := workflow.SideEffect(ctx, func(ctx workflow.Context) interface{} {
			return rand.Intn(max)
		})
		err := encodedNUmber.Get(&number)
		if err != nil {
			workflow.GetLogger(ctx).Info("error", err)
			return Output{}, err
		}
	}
	for j := 1; j <= number; j++ {
		// Don't use GO's time functions because they are not deterministic
		workflow.Sleep(ctx, 1*time.Second)
		r *= j
	}
	return Output{Result: r}, nil
}
