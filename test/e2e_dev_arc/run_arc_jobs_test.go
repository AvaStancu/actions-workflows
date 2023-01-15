package e2e_dev_arc

import (
	"flag"
	"fmt"
	"testing"

	"strings"
)

var availablePods = flag.String("pods", "", "Should receive the output of a kubectl get pods command.")

func TestARCJobs(t *testing.T) {
	flag.Parse()

	expectedPods := []string{"listener", "runner", "controller-manager"}

	t.Run("Get available pods after job run", func(t *testing.T) {
		fmt.Printf("Here we go")
		for _, podName := range expectedPods {
			if !strings.Contains(*availablePods, podName) {
				t.Fatalf("%v pod not found.", podName)
			}
		}
	},
	)
}
