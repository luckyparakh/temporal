package example05

import (
	"go.temporal.io/sdk/workflow"
)

// Code in workflow should be deterministic
func Workflow(ctx workflow.Context) error {
	workflow.GetLogger(ctx).Info("starting example 05")
	var number int
	// docker exec temporal-admin-tools tctl wf query --wid 157fecbb-dfab-4246-b1b9-75ec8aea8ef7 -qt current_number
	err := workflow.SetQueryHandler(ctx, "current_number", func() (int, error) {
		return number, nil
	})
	if err != nil {
		return err
	}
	// Don't use the GO channels (select also) because these are not deterministic
	// Use temporal's wrapper instead
	// docker exec temporal-admin-tools tctl wf signal --wid 157fecbb-dfab-4246-b1b9-75ec8aea8ef7 -n set_number -i 69
	signalChan := workflow.GetSignalChannel(ctx, "set_number")

	s := workflow.NewSelector(ctx)
	s.AddReceive(signalChan, func(c workflow.ReceiveChannel, more bool) {
		c.Receive(ctx, &number)
		workflow.GetLogger(ctx).Info("Received new number", number)
	})
	for {
		s.Select(ctx)
	}
}
