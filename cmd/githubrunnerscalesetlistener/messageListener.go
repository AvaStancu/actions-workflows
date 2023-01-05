package main

import (
	"context"

	"github.com/actions-runner-controller/actions-runner-controller/github/actions"
)

//go:generate mockery --inpackage --name=RunnerScaleSetClient
type RunnerScaleSetClient interface {
	GetRunnerScaleSetMessage(ctx context.Context, handler func(msg *actions.RunnerScaleSetMessage) error) error
	AcquireJobsForRunnerScaleSet(ctx context.Context, requestIds []int64) error
}
