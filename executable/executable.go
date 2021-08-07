package executable

import (
	"fmt"
	"github.com/NubeIO/nube-rubix-updater-cli/cli"
)

// Execute configures the CLI and executes the root command
func Execute() {

	err := cli.ConfigureFromFile("cli.yml")
	if err != nil {
		fmt.Println(err)
	}
	cli.Execute()
}
