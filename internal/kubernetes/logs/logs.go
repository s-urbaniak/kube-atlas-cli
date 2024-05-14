package logs

import (
	"compress/gzip"
	"context"
	"fmt"
	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1"
	"github.com/s-urbaniak/kube-atlas-cli/internal/store"
	"go.mongodb.org/atlas-sdk/v20231115013/admin"
	"golang.org/x/sync/errgroup"
	"io"
	"net/url"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"strings"
	"time"
)

type Logs struct {
	namespace           string
	atlasDeploymentName string
	logName             string
	hostname            string
	k8sClient           client.Client
	logsDownloader      store.LogsDownloader
	follow              bool
}

func (l *Logs) WithHostname(hostname string) *Logs {
	l.hostname = hostname
	return l
}

func (l *Logs) WithLogName(logName string) *Logs {
	l.logName = logName
	return l
}

func (l *Logs) WithNamespace(ns string) *Logs {
	l.namespace = ns
	return l
}

func (l *Logs) WithAtlasDeploymentName(atlasDeploymentName string) *Logs {
	l.atlasDeploymentName = atlasDeploymentName
	return l
}

func (l *Logs) WithLogsDownloader(logsDownloader store.LogsDownloader) *Logs {
	l.logsDownloader = logsDownloader
	return l
}

func (l *Logs) WithK8sClient(client client.Client) *Logs {
	l.k8sClient = client
	return l
}

func (l *Logs) WithFollow(follow bool) *Logs {
	l.follow = follow
	return l
}

func (l *Logs) Run(ctx context.Context) error {
	deployment := akov2.AtlasDeployment{}
	err := l.k8sClient.Get(ctx, client.ObjectKey{Name: l.atlasDeploymentName, Namespace: l.namespace}, &deployment)
	if err != nil {
		return fmt.Errorf("error getting AtlasDeployment: %w", err)
	}

	possibleHostnames := "N/A"
	if deployment.Status.ConnectionStrings != nil {
		u, err := url.Parse(deployment.Status.ConnectionStrings.Standard)
		if err != nil {
			return fmt.Errorf("error parsing connection strings: %w", err)
		}
		hosts := strings.Split(u.Host, ",")
		for i, h := range hosts {
			hosts[i] = strings.Split(h, ":")[0]
		}
		possibleHostnames = strings.Join(hosts, ", ")
	}

	if l.hostname == "" {
		return fmt.Errorf("hostname not specified, possible values are: %v", possibleHostnames)
	}

	project := akov2.AtlasProject{}
	err = l.k8sClient.Get(ctx, deployment.AtlasProjectObjectKey(), &project)
	if err != nil {
		return fmt.Errorf("error getting AtlasProject: %w", err)
	}

	// as per https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Monitoring-and-Logs,
	// poll every 5 minutes.
	ticker := time.NewTicker(5 * time.Minute)

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		for {
			// starting anonymous function in order to defer closing the reader at each iteration.
			if err := func() error {
				params := &admin.GetHostLogsApiParams{
					GroupId:  project.Status.ID,
					HostName: l.hostname,
					LogName:  l.logName,
				}
				reader, err := l.logsDownloader.DownloadLog(params)
				if err != nil {
					return fmt.Errorf("error downloading logs: %w", err)
				}
				defer reader.Close()

				gzipReader, err := gzip.NewReader(reader)
				if err != nil {
					return fmt.Errorf("error creating gzip reader: %w", err)
				}

				_, err = io.Copy(os.Stdout, gzipReader)
				if err != nil {
					return fmt.Errorf("error reading logs: %w", err)
				}

				return nil
			}(); err != nil {
				return err
			}

			// bail out immediately, if we are not following
			if !l.follow {
				return nil
			}

			select {
			case <-ctx.Done():
				fmt.Fprintf(os.Stderr, "aborting log streaming\n")
				return nil
			case <-ticker.C:
			}
		}
	})
	return g.Wait()
}

func NewLogs() *Logs {
	return &Logs{}
}
