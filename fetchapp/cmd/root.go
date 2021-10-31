package cmd

import (
	"os"

	"github.com/marprin/assessment/fetchapp/cmd/http"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "fetchapp",
		Short: "Fetchapp - Backend Service",
	}
)

func init() {
	rootCmd.AddCommand(http.ServeHTTPCmd())
}

// Execute is the exposed function
func Execute() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := rootCmd.Execute(); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("Error when initiate fetchapp service up")
		os.Exit(1)
	}
}
