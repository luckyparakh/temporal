package example06

import "context"

type Input struct {
	A int
	B int
}

type Output struct {
	C int
}


// Activity can use GOlang Context 
func Activity06(ctx context.Context,input Input) (Output, error) {
	return Output{input.A + input.B}, nil
}
