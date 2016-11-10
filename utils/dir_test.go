package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"path/filepath"
	"testing"
)

type UtilsDirTestSuite struct {
	suite.Suite
}

func (suite *UtilsDirTestSuite) SetupTest() {
	path := filepath.Join(tmpPath...)
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		if err := os.RemoveAll(path); err != nil {
			fmt.Println(err)
		}
	}
}

func (suite *UtilsDirTestSuite) TeardownTest() {
	dir := filepath.Join(tmpPath...)
	if err := os.RemoveAll(dir); err != nil {
		fmt.Println(err)
	}
}
func (suite *UtilsDirTestSuite) TestMkTmpdir() {
	CreateTmpDir()
	path := filepath.Join(tmpPath...)
	_, err := os.Stat(path)
	assert.True(suite.T(), !os.IsNotExist(err))
}

func TestUtilsDirTestSuite(t *testing.T) {
	suite.Run(t, new(UtilsDirTestSuite))
}
