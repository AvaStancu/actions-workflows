package e2e_dev_arc

import (
	"flag"
	"fmt"
	"testing"

	"golang.org/x/exp/slices"
)

func TestARCJobs(t *testing.T) {
	availablePods := flag.String("pods", "", "Should receive the output of a kubectl get pods command.")
	flag.Parse()

	expectedPods := []string{"listener", "runner", "controller-manager"}

	t.Run("Get available pods after job run", func(t *testing.T) {
		fmt.Printf("Here we go")
		for _, podName := range expectedPods {
			if !slices.Contains(*availablePods, podName) {
				t.Fatalf("%v pod not found.", podName)
			}
		}
	},
	)
}
