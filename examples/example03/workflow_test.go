package example03

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"go.temporal.io/sdk/testsuite"
)

func TestWorkflowTestSuite(t *testing.T) {
	suite.Run(t, new(WorkflowTestSuite))
}

type WorkflowTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite
	// Consider it as an Worker
	env *testsuite.TestWorkflowEnvironment
}

func (s *WorkflowTestSuite) SetupTest() {
	s.env = s.NewTestWorkflowEnvironment()
}

func (s *WorkflowTestSuite) AfterTest(suiteName, testName string) {
	s.env.AssertExpectations(s.T())
}

func (s *WorkflowTestSuite) Test_WF() {
	s.env.RegisterWorkflow(Workflow)
	s.env.ExecuteWorkflow(Workflow, Input{1, 2})
	s.Require().True(s.env.IsWorkflowCompleted())
	s.Require().NoError(s.env.GetWorkflowError())
	var o Output
	s.Require().NoError(s.env.GetWorkflowResult(&o))
	s.Equal(3, o.Result)
}
