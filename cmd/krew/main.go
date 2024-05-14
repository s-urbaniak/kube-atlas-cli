package main

import (
	"context"
	"fmt"
	"github.com/rajatjindal/krew-release-bot/pkg/releaser"
	"github.com/rajatjindal/krew-release-bot/pkg/source"
	"os"
)

func main() {
	err := process(context.Background())
	if err != nil {
		fmt.Printf("\n%v\n", err)
		os.Exit(1)
	}
}

func process(ctx context.Context) error {
	releaseRequest := &source.ReleaseRequest{
		TagName:            "v0.0.1",
		PluginName:         "atlas",
		PluginRepo:         "kube-atlias-cli",
		PluginOwner:        "s-urbaniak",
		PluginReleaseActor: "s-urbaniak",
		TemplateFile:       ".krew.yaml",
	}

	pluginName, pluginManifest, err := source.ProcessTemplate(".krew.yaml", releaseRequest)
	if err != nil {
		return fmt.Errorf("error processing template: %w", err)
	}

	releaseRequest.PluginName = pluginName
	releaseRequest.ProcessedTemplate = pluginManifest

	fmt.Printf("releasing %q, manifest:\n%s\n", pluginName, pluginManifest)

	r := releaser.New(os.Getenv("GH_TOKEN"))
	r.TokenEmail = "sergiusz.urbaniak@gmail.com"
	r.TokenUserHandle = "s-urbaniak"
	r.TokenUsername = "Sergiusz Urbaniak"
	r.LocalKrewIndexRepoCloneURL = r.UpstreamKrewIndexRepoCloneURL
	r.LocalKrewIndexRepo = r.UpstreamKrewIndexRepo
	r.LocalKrewIndexRepoOwner = r.UpstreamKrewIndexRepoOwner

	pr, err := r.Release(releaseRequest)
	if err != nil {
		return fmt.Errorf("error releasing request: %w", err)
	}
	fmt.Println(pr)
	return nil
}
