package cmd

import (
	"fmt"

	ver "github.com/nao1215/spare/version"
	"github.com/spf13/cobra"
)

// newVersionCmd return version command.
func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show " + ver.Name + " command version information",
		Run:   version,
	}
}

// version return spare command version.
func version(_ *cobra.Command, _ []string) {
	fmt.Printf("%s version %s, revision %s (under MIT LICENSE)\n", ver.Name, ver.TagVersion, ver.Revision)
}
