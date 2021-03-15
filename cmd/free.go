package cmd

import (
	"fmt"

	"github.com/jaedle/golang-tplink-hs100/pkg/configuration"
	"github.com/jaedle/golang-tplink-hs100/pkg/hs100"
	"github.com/spf13/cobra"
)

var (
	freeCmd = &cobra.Command{
		Use:   "free",
		Short: "Turns outlet off to indicate free",
		Long:  `Turns outlet off to indicate free`,
		RunE: func(cmd *cobra.Command, args []string) error {
			outlet = hs100.NewHs100(address, configuration.Default())
			err := outlet.TurnOff()
			if err != nil {
				return fmt.Errorf("Error turning outlet on: %s", err)
			}
			return nil
		},
	}
)
