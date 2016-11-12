package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/timakin/pppb/constants"
)

func CreateTmpDir() {
	dir := filepath.Join(constants.tmpPath...)
	if err := os.MkdirAll(dir, 0777); err != nil {
		fmt.Println(err)
	}
}
