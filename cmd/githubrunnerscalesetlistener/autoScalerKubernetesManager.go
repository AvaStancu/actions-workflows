package main

import (
	"context"
	"fmt"

	"github.com/actions-runner-controller/actions-runner-controller/api/v1alpha1"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type AutoScalerKubernetesManager struct {
	*kubernetes.Clientset

	logger logr.Logger
}

func NewKubernetesManager(logger *logr.Logger) (*AutoScalerKubernetesManager, error) {
	conf, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	kubeClient, err := kubernetes.NewForConfig(conf)
	if err != nil {
		return nil, err
	}

	var manager = &AutoScalerKubernetesManager{
		Clientset: kubeClient,
		logger:    logger.WithName("KubernetesManager"),
	}
	return manager, nil
}

func (k *AutoScalerKubernetesManager) ScaleEphemeralRunnerSet(ctx context.Context, namespace, resourceName string, runnerCount int) error {
	patchedEphemeralRunnerSet := &v1alpha1.EphemeralRunnerSet{}
	patchJson := fmt.Sprintf("{\"spec\":{\"replicas\":%d}}", runnerCount)
	err := k.RESTClient().
		Patch(types.MergePatchType).
		Prefix("apis", "actions.summerwind.dev", "v1alpha1").
		Namespace(namespace).
		Resource("EphemeralRunnerSets").
		Name(resourceName).
		Body([]byte(patchJson)).
		Do(ctx).
		Into(patchedEphemeralRunnerSet)
	if err != nil {
		return fmt.Errorf("could not patch ephemeral runner set , patch JSON: %s, error: %w", string(patchJson), err)
	}

	k.logger.Info("Ephemeral runner set scaled.", "namespace", namespace, "name", resourceName, "replicas", patchedEphemeralRunnerSet.Spec.Replicas)
	return nil
}
