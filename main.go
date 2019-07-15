package main

import (
	"fmt"
	"github.com/jumale/alfred-go/cmd"
	"github.com/jumale/alfred-go/pkg/k8s"
	"os"
)

func main() {
	app := cmd.NewRootCmd(cmd.ModulesConfig{
		K8s: &k8s.Config{
			BinPath: "kubectl",
		},
	})
	if err := app.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
