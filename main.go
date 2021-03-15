package main

import (
	"os"

	"github.com/jaedle/golang-tplink-hs100/pkg/configuration"
	"github.com/jaedle/golang-tplink-hs100/pkg/hs100"
)

func main() {

	h := hs100.NewHs100(os.Getenv(("OUTLET_IP")), configuration.Default())

	h.TurnOn()
	h.TurnOff()
}
