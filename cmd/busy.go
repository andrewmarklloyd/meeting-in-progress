package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	busyCmd = &cobra.Command{
		Use:   "busy",
		Short: "Turns outlet on to indicate busy",
		Long:  `Turns outlet on to indicate busy`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := outlet.TurnOn()
			if err != nil {
				return fmt.Errorf("Error turning outlet on: %s", err)
			}
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(busyCmd)
}
