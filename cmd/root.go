package cmd

import (
	"os"

	"github.com/jaedle/golang-tplink-hs100/pkg/configuration"
	"github.com/jaedle/golang-tplink-hs100/pkg/hs100"
	"github.com/spf13/cobra"
)

var (
	outlet *hs100.Hs100

	rootCmd = &cobra.Command{
		Use:   "meeting-in-progress",
		Short: "Triggers light indicating meeting status",
		Long: `Working from home means possible interruptions
from people in your house. Use this tool to trigger a
smart outlet indicating you are busy and should 
not be interrupted.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	outlet = hs100.NewHs100(os.Getenv(("OUTLET_IP")), configuration.Default())
}
