package main

import (
	"context"
	"fmt"
	"github.com/rajatjindal/krew-release-bot/pkg/releaser"
	"github.com/rajatjindal/krew-release-bot/pkg/source"
	"os"
)

func main() {
	pr, err := process(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "\n%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n%v\n", pr)
}

func process(ctx context.Context) (string, error) {
	releaseRequest := &source.ReleaseRequest{
		TagName:            os.Getenv("TAG_NAME"),
		PluginName:         "atlas",
		PluginRepo:         "kube-atlias-cli",
		PluginOwner:        "s-urbaniak",
		PluginReleaseActor: "s-urbaniak",
		TemplateFile:       ".krew.yaml",
	}

	pluginName, pluginManifest, err := source.ProcessTemplate(".krew.yaml", releaseRequest)
	if err != nil {
		return "", fmt.Errorf("error processing template: %w", err)
	}

	releaseRequest.PluginName = pluginName
	releaseRequest.ProcessedTemplate = pluginManifest

	fmt.Printf("releasing %q, manifest:\n%s\n", pluginName, pluginManifest)

	r := releaser.New(os.Getenv("GITHUB_TOKEN"))
	r.TokenEmail = "sergiusz.urbaniak@gmail.com"
	r.TokenUserHandle = "s-urbaniak"
	r.TokenUsername = "Sergiusz Urbaniak"
	r.LocalKrewIndexRepoCloneURL = r.UpstreamKrewIndexRepoCloneURL
	r.LocalKrewIndexRepo = r.UpstreamKrewIndexRepo
	r.LocalKrewIndexRepoOwner = r.UpstreamKrewIndexRepoOwner

	pr, err := r.Release(releaseRequest)
	if err != nil {
		return "", fmt.Errorf("error releasing request: %w", err)
	}
	return pr, nil
}
