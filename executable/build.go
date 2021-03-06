package executable

import (
	"fmt"
	"github.com/NubeIO/nube-rubix-updater-cli/cli"
	"io/ioutil"
	"log"
	"os"
	"runtime"

	ex "github.com/egjiri/go-kit/exec"
	yaml "gopkg.in/yaml.v2"
)

type config struct {
	Name string
}

func init() {
	cli.AddRunToCommand("build", func(c *cli.Command) {
		currentPath, err := os.Getwd()
		if err != nil {
			log.Fatal("Error: ", err)
		}

		goos := c.FlagString("goos")
		if goos == "" {
			goos = runtime.GOOS
		}
		goarch := c.FlagString("goarch")
		if goarch == "" {
			goarch = runtime.GOARCH
		}
		gobuildFlags := c.FlagString("gobuild-flags")
		// TODO: Figure out best way of versioning the docker image instead of defaulting to latest
		command := fmt.Sprintf("docker run --rm -v %s:/data -e GOOS_TARGET=%s -e GOARCH_TARGET=%s -e GOBUILD_FLAGS=\"%s\" -e REPO=%s egjiri/cliff", currentPath, goos, goarch, gobuildFlags, c.Arg(0))
		if err := ex.Execute(command); err != nil {
			log.Fatal("Error: ", err, "\n", command)
		}

		newName := fmt.Sprintf("%s/%s", c.FlagString("output"), name())
		if err := os.Rename("cliff-binary", newName); err != nil {
			log.Fatal("Error: ", err)
		}
		fmt.Println("Built binary:", newName)
	})
}

func name() string {
	yamlConfigContent, err := ioutil.ReadFile("cli.yml")
	if err != nil {
		log.Fatal(err)
	}

	var c config
	if err := yaml.Unmarshal(yamlConfigContent, &c); err != nil {
		log.Fatal(err)
	}
	return c.Name
}
