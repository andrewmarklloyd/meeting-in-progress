package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	freeCmd = &cobra.Command{
		Use:   "free",
		Short: "Turns outlet off to indicate free",
		Long:  `Turns outlet off to indicate free`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := outlet.TurnOff()
			if err != nil {
				return fmt.Errorf("Error turning outlet on: %s", err)
			}
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(freeCmd)
}
