package cmd

import (
	"fmt"
	"os"

	"github.com/jaedle/golang-tplink-hs100/pkg/hs100"
	"github.com/spf13/cobra"
)

var (
	outlet  *hs100.Hs100
	address string
	rootCmd = &cobra.Command{
		Use:   "meeting-indicator",
		Short: "Triggers light indicating meeting status",
		Long: `Working from home means possible interruptions
from people in your house. Use this tool to trigger a
smart outlet indicating you are busy and should
not be interrupted.`,
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&address, "address", "a", "", "Address or IP of the smart outlet")
	rootCmd.MarkPersistentFlagRequired("address")
	rootCmd.AddCommand(freeCmd)
	rootCmd.AddCommand(busyCmd)
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
