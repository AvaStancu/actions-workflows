package e2e_dev_arc

import (
	// "flag"
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// var availablePods = flag.String("pods", "", "Should receive the output of a kubectl get pods command.")

type podCountsByType struct {
	controllers int
	listeners   int
	runners     int
}

func getPodsByType(clientset *kubernetes.Clientset) podCountsByType {
	namespace := "actions-runner-system"
	availablePods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	podsByType := podCountsByType{}
	for _, pod := range availablePods.Items {
		if strings.Contains(pod.Name, "runner") {
			podsByType.runners += 1
		}
		if strings.Contains(pod.Name, "controller-manager") {
			podsByType.controllers += 1
		}
		if strings.Contains(pod.Name, "listener") {
			podsByType.listeners += 1
		}
	}
	return podsByType
}

func pollForClusterState(clientset *kubernetes.Clientset, expectedPodsCount podCountsByType, maxTime int) bool {
	sleepTime := 5
	maxRetries := maxTime / sleepTime
	success := false
	for i := 0; i <= maxRetries; i++ {
		time.Sleep(time.Second * time.Duration(sleepTime))
		availablePodsCount := getPodsByType(clientset)
		if availablePodsCount == expectedPodsCount {
			success = true
			break
		} else {
			fmt.Printf("%v", availablePodsCount)
		}
	}
	return success
}

func TestARCJobs(t *testing.T) {
	// flag.Parse()

	configFile := filepath.Join(
		os.Getenv("HOME"), ".kube", "config",
	)

	config, err := clientcmd.BuildConfigFromFlags("", configFile)
	if err != nil {
		t.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Get available pods before job run", func(t *testing.T) {
		expectedPodsCount := podCountsByType{1, 1, 1}
		success := pollForClusterState(clientset, expectedPodsCount, 20)
		if !success {
			t.Fatal("Expected pods count did not match available pods count before job run.")
		}
	},
	)
	t.Run("Get available pods during job run", func(t *testing.T) {
		c := http.Client{}
		url := "https://api.github.com/repos/AvaStancu/actions-workflows/actions/workflows/44661067/dispatches"
		var jsonStr = []byte(`{"ref":"main"}`)

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		if err != nil {
			t.Fatal(err)
		}
		ght := os.Getenv("GITHUB_TOKEN")
		req.Header.Add("Accept", "application/vnd.github+json")
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ght))
		req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

		resp, err := c.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		expectedPodsCount := podCountsByType{1, 1, 3}
		success := pollForClusterState(clientset, expectedPodsCount, 60)
		if !success {
			t.Fatal("Expected pods count did not match available pods count during job run.")
		}

	},
	)
	t.Run("Get available pods after job run", func(t *testing.T) {
		expectedPodsCount := podCountsByType{1, 1, 1}
		success := pollForClusterState(clientset, expectedPodsCount, 120)
		if !success {
			t.Fatal("Expected pods count did not match available pods count after job run.")
		}
	},
	)
}
