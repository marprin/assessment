package http

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	serveHTTPCmd = &cobra.Command{
		Use:   "serve",
		Short: "Fetch App HTTP Service",
		RunE:  run,
	}
)

func run(cmd *cobra.Command, args []string) error {
	logrus.Infoln("Starting initialize http service")

	logrus.Infoln("Terminate http service")

	return nil
}

// ServeHTTPCmd is the exposed function
func ServeHTTPCmd() *cobra.Command {
	return serveHTTPCmd
}
