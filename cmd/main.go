package main

import (
	"context"
	"os"

	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
	"github.com/tuihub/tuihub-go"
	"github.com/tuihub/tuihub-go/logger"
	"github.com/tuihub/tuihub-telegram/internal"
)

// go build -ldflags "-X main.version=x.y.z".
var (
	// version is the version of the compiled software.
	version string
)

func main() {
	config := tuihub.PorterConfig{
		Name:       "tuihub-telegram",
		Version:    version,
		GlobalName: "github.com/tuihub/tuihub-telegram",
		FeatureSummary: &porter.PorterFeatureSummary{
			SupportedAccounts:    nil,
			SupportedAppSources:  nil,
			SupportedFeedSources: nil,
			SupportedNotifyDestinations: []string{
				tuihub.WellKnownToString(
					librarian.WellKnownNotifyDestination_WELL_KNOWN_NOTIFY_DESTINATION_TELEGRAM,
				),
			},
		},
	}
	server, err := tuihub.NewPorter(
		context.Background(),
		config,
		internal.NewHandler(),
	)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	if err = server.Run(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
