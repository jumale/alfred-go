package cmd

import (
	"github.com/jumale/alfred-go/pkg/k8s"
	"github.com/spf13/cobra"
)

type ModulesConfig struct {
	K8s *k8s.Config
}

func NewRootCmd(cfg ModulesConfig) *cobra.Command {
	root := &cobra.Command{}
	if cfg.K8s != nil {
		root.AddCommand(newK8sCommand(*cfg.K8s))
	}

	return root
}
