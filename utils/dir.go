package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

var tmpPath = []string{"/tmp", "pppb_tmp"}

func CreateTmpDir() {
	dir := filepath.Join(tmpPath...)
	if err := os.MkdirAll(dir, 0777); err != nil {
		fmt.Println(err)
	}
}
