package main

import (
	"context"
)

//go:generate mockery --inpackage --name=KubernetesManager
type KubernetesManager interface {
	ScaleEphemeralRunnerSet(ctx context.Context, namespace, resourceName string, runnerCount int) error
}
