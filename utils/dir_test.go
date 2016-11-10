package utils

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestMkTmpdir(t *testing.T) {
	utils.CreateTmpDir()
	path := filepath.Join(utils.tmpPath)
	assert.True(t, os.IsExist(path))
}
