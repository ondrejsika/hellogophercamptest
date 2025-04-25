package root

import (
	"github.com/ondrejsika/hellogophercamptest/version"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "hellogophercamptest",
	Short: "hellogophercamptest, " + version.Version,
}
