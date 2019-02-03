package main

import (
	"github.com/pwoelfle/synology-cli/cli/cmd"
	_ "github.com/pwoelfle/synology-cli/cli/cmd/downloadstation"
	_ "github.com/pwoelfle/synology-cli/cli/cmd/downloadstation/task"
	_ "github.com/pwoelfle/synology-cli/cli/cmd/downloadstation/task/add"
	_ "github.com/pwoelfle/synology-cli/cli/cmd/downloadstation/task/describe"
	_ "github.com/pwoelfle/synology-cli/cli/cmd/downloadstation/task/list"
)

func main() {
	cmd.Execute()
}
