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
		Run: func(cmd *cobra.Command, args []string) {
			err := outlet.TurnOff()
			if err != nil {
				fmt.Println(fmt.Errorf("Error turning outlet off: %s", err))
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(busyCmd)
}
