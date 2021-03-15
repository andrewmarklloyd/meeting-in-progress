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
		Run: func(cmd *cobra.Command, args []string) {
			err := outlet.TurnOff()
			if err != nil {
				fmt.Println(fmt.Errorf("Error turning outlet off: %s", err))
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(freeCmd)
}
