package logs

import (
	"context"
	"fmt"
	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1"
	"github.com/s-urbaniak/kube-atlas-cli/internal/cli"
	"github.com/s-urbaniak/kube-atlas-cli/internal/config"
	"github.com/s-urbaniak/kube-atlas-cli/internal/flag"
	"github.com/s-urbaniak/kube-atlas-cli/internal/kubernetes/logs"
	"github.com/s-urbaniak/kube-atlas-cli/internal/pointer"
	"github.com/s-urbaniak/kube-atlas-cli/internal/prerun"
	"github.com/s-urbaniak/kube-atlas-cli/internal/store"
	"github.com/s-urbaniak/kube-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type LogsOpts struct {
	cli.GlobalOpts
	*genericclioptions.ConfigFlags
	logname  string
	hostname string
	follow   bool
	store    *store.Store
}

func (opts *LogsOpts) Run(ctx context.Context, deploymentName string) error {
	restConfig, err := opts.ToRESTConfig()
	if err != nil {
		return fmt.Errorf("failed to create REST config: %w", err)
	}

	err = akov2.AddToScheme(scheme.Scheme)
	if err != nil {
		return fmt.Errorf("error registering scheme: %w", err)
	}

	k8sClient, err := client.New(restConfig, client.Options{Scheme: scheme.Scheme})
	if err != nil {
		return fmt.Errorf("unable to setup kubernetes client: %w", err)
	}

	ns := opts.Namespace
	if ns == nil || *ns == "" {
		ns = pointer.Get("default")
	}

	err = logs.NewLogs().
		WithNamespace(*ns).
		WithAtlasDeploymentName(deploymentName).
		WithK8sClient(k8sClient).
		WithLogname(opts.logname).
		WithHostname(opts.hostname).
		WithFollow(opts.follow).
		WithLogsDownloader(opts.store).
		Run(ctx)

	if err != nil {
		return fmt.Errorf("error running logs command: %w", err)
	}

	return nil
}

func (opts *LogsOpts) initStores(ctx context.Context) prerun.CmdOpt {
	return func() error {
		var err error

		profile := config.Default()
		opts.store, err = store.New(store.AuthenticatedPreset(profile), store.WithContext(ctx))
		if err != nil {
			return err
		}

		return nil
	}
}

func Builder() *cobra.Command {
	const use = "logs"
	opts := &LogsOpts{ConfigFlags: genericclioptions.NewConfigFlags(false)}
	cmd := &cobra.Command{
		Use:   use,
		Short: "Deployment logs",
		Args:  cobra.ExactArgs(1),
		Long:  `Stream logs that contains a range of log messages for the specified deployment for the specified project.`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initStores(cmd.Context()),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), args[0])
		},
	}

	cmd.Flags().BoolVarP(&opts.follow, "follow", "f", false, `Specify if the logs should be streamed.`)
	cmd.Flags().StringVarP(&opts.hostname, "hostname", "H", "", `Human-readable label that identifies the host that stores the log files that you want to download.`)
	cmd.Flags().StringVarP(&opts.logname, "logname", "l", "mongodb", `Human-readable label that identifies the log file that you want to return. 
To return audit logs, enable Database Auditing for the specified project.
Possible values are "mongodb" "mongos" "mongodb-audit-log" "mongos-audit-log".`)
	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)
	opts.ConfigFlags.AddFlags(cmd.Flags())

	return cmd
}
