package cmd

import (
	"fmt"

	"github.com/jaedle/golang-tplink-hs100/pkg/configuration"
	"github.com/jaedle/golang-tplink-hs100/pkg/hs100"
	"github.com/spf13/cobra"
)

var (
	busyCmd = &cobra.Command{
		Use:   "busy",
		Short: "Turns outlet on to indicate busy",
		Long:  `Turns outlet on to indicate busy`,
		RunE: func(cmd *cobra.Command, args []string) error {
			outlet = hs100.NewHs100(address, configuration.Default())
			err := outlet.TurnOn()
			if err != nil {
				return fmt.Errorf("Error turning outlet on: %s", err)
			}
			return nil
		},
	}
)
