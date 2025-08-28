package workflow

import (
	"context"

	"go.temporal.io/sdk/activity"
)

func CreateVclusterActivity(ctx context.Context, input VclusterInput) error {
	return triggerGithubAction(ctx, "create", input)
}

func SignalCompleteActivity(ctx context.Context, input VclusterInput) error {
	activity.GetLogger(ctx).Info("Vcluster creation process is complete.")
	return nil
}
