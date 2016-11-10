package command

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"path/filepath"
)

var tmpPath = []string{"/tmp", "pppb_tmp"}

func CmdInspect(c *cli.Context) {
	createTmpDir()
}

func createTmpDir() {
	dir := filepath.Join(tmpPath...)

	if err := os.Mkdir(dir, 0777); err != nil {
		fmt.Println(err)
	}
}
