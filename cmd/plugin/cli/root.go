package cli

import (
	"fmt"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/cli-runtime/pkg/genericclioptions"

	configcmd "github.com/s-urbaniak/kube-atlas-cli/cmd/plugin/cli/config"
	logscmd "github.com/s-urbaniak/kube-atlas-cli/cmd/plugin/cli/logs"
	operatorcmd "github.com/s-urbaniak/kube-atlas-cli/cmd/plugin/cli/operator"
	"github.com/s-urbaniak/kube-atlas-cli/internal/config"
)

var (
	KubernetesConfigFlags *genericclioptions.ConfigFlags
)

func RootCmd() *cobra.Command {
	const use = "atlas"
	cmd := &cobra.Command{
		Use:   use,
		Short: "Manage Kubernetes resources.",
		Long:  `This command provides access to Kubernetes features within Atlas.`,
	}

	cobra.OnInitialize(initConfig)

	k8sFlags := genericclioptions.NewConfigFlags(false)
	k8sFlags.AddFlags(cmd.PersistentFlags())
	cmd.SetContext(signals.SetupSignalHandler()) // wire SIGTERM and SIGINT
	cmd.AddCommand(configcmd.Builder(), operatorcmd.Builder(), logscmd.Builder(k8sFlags))
	return cmd
}

func InitAndExecute() {
	if err := loadConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := RootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func loadConfig() error {
	if err := config.LoadAtlasCLIConfig(); err != nil {
		return fmt.Errorf("error loading config: %w. Please run `atlas config init` to reconfigure your profile", err)
	}

	return nil
}

func initConfig() {
	viper.AutomaticEnv()
}
