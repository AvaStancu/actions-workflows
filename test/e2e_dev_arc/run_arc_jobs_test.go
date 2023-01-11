package e2e_dev_arc

import ("flag", "testing")

import  "golang.org/x/exp/slices"


func TestARCJobs(t *testing.T){
    availablePods := flag.String("pods")
    flag.Parse()

    expectedPods = []string{"listener", "runner", "controller-manager"}

    t.Run("Get available pods after job run", func(t *testing.T) {
      for _, podName in range := expectedPods {
          if !slices.Contains(*availablePods, podName) {
              t.Fatalf("%v pod not found.", podName)
            }
        }
    }
}
