package cmd

import (
	_ "github.com/ondrejsika/hellogophercamptest/cmd/gopher"
	"github.com/ondrejsika/hellogophercamptest/cmd/root"
	_ "github.com/ondrejsika/hellogophercamptest/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
