package cmd

import (
	"encoding/json"
	"github.com/jumale/alfred-go/pkg/k8s"
	"github.com/jumale/alfred-go/pkg/k8s/cli"
	"github.com/spf13/cobra"
	"os"
	"text/template"
)

func newK8sCommand(cfg k8s.Config) *cobra.Command {
	if cfg.PodTitle == nil {
		cfg.PodTitle = k8s.DefaultTitle
	}
	if cfg.ContainerTitle == nil {
		cfg.ContainerTitle = k8s.DefaultTitle
	}
	if cfg.NamespaceTitle == nil {
		cfg.NamespaceTitle = k8s.DefaultTitle
	}
	if cfg.ContextTitle == nil {
		cfg.ContextTitle = k8s.DefaultTitle
	}

	if cfg.ContextOrder == nil {
		cfg.ContextOrder = k8s.DefaultOrder
	}

	podBashCmd := cli.CmdPodExec("{{.Pod}}", "{{.Container}}", "-it", "bash")
	podBashTpl := template.Must(template.New("pod_exec").Parse(podBashCmd))

	k8sClient := &cli.Client{BinPath: cfg.BinPath}
	pods := &k8s.PodsFilter{Client: k8sClient, Config: cfg}
	//contexts := &k8s.ContextsFilter{Client: k8sClient, Config: cfg}
	//namespaces := &k8s.NamespacesFilter{Client: k8sClient, Config: cfg}
	root := &cobra.Command{
		Use:   "k8s",
		Short: "k8s commands",
	}

	root.AddCommand(&cobra.Command{
		Use:   "bash",
		Short: "Exec bash to a pod",
		RunE: func(cmd *cobra.Command, args []string) error {
			query := ""
			if len(args) == 1 {
				query = args[0]
			}
			list, err := pods.Filter(query, podBashTpl)
			if err != nil {
				return err
			}
			return json.NewEncoder(os.Stdout).Encode(list)
		},
	})

	return root
}

func getQuery() {

}
