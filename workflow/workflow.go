package workflow

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

type VclusterInput struct {
	VclusterName string
	Namespace    string
	CPU          string
	Memory       string
	Storage      string

	WorkflowID string
}

func CreateVclusterWorkflow(ctx workflow.Context, input VclusterInput) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	if err := workflow.ExecuteActivity(ctx, CreateVclusterActivity, input).Get(ctx, nil); err != nil {
		workflow.GetLogger(ctx).Error("CreateVclusterActivity failed", "error", err)
		return "", err
	}

	// The workflow now completes here.

	return "vCluster creation workflow triggered successfully", nil
}
