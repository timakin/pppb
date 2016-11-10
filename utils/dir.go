package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

var tmpPath = []string{"/tmp", "pppb_tmp"}

func CreateTmpDir() {
	dir := filepath.Join(tmpPath...)
	_, err := os.Stat(dir)
	if os.IsExist(err) {
		return
	}
	if err := os.Mkdir(dir, 0777); err != nil {
		fmt.Println(err)
	}
}
