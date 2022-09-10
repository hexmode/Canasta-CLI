package version

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/CanastaWiki/Canasta-CLI-Go/internal/config"
)

var (
	sha1 string
	buildTime string
)

var instance config.Installation

func NewCmdCreate() *cobra.Command {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Show the Canasta version",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf( "This is canasta: built at %s from git commit %s.\n", buildTime, sha1 )
			os.Exit(0)
			return nil
		},
	}
	return versionCmd
}
