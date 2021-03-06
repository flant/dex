package main

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"

	"github.com/dexidp/dex/v2/version"
)

func commandVersion() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version and exit",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf(
				"Dex Version: %s\nGo Version: %s\nGo OS/ARCH: %s %s\n",
				version.Version,
				runtime.Version(),
				runtime.GOOS,
				runtime.GOARCH,
			)
		},
	}
}
