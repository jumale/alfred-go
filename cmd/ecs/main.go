package main

import "github.com/jumale/alfred-go/pkg/k8s/kubectl"

func main() {
	k8s := kubectl.NewClient(kubectl.Config{})

	k8s.Pods()
}
