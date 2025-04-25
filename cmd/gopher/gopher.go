package gopher

import (
	"fmt"

	gopher "github.com/ondrejsika/go-gopher"
	"github.com/ondrejsika/hellogophercamptest/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "gopher",
	Aliases: []string{"g"},
	Short:   "Prints Gopher to CLI",
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Print(gopher.GOPHER)
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
}
