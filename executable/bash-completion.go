package executable

import "github.com/NubeIO/nube-rubix-updater-cli/cli"

func init() {
	cli.AddRunToCommand("bash-completion", func(c *cli.Command) {
		outputPath := c.FlagString("output")
		cli.ConfigureAndGenerateBashCompletionFile(outputPath)
	})
}
