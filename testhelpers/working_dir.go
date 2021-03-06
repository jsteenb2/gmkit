package testhelpers

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// WorkDir creates a new temporary working directory while storing
// the current working directory for restoration at a later time
type WorkDir struct {
	cwd string
	tmp string
}

// NewWorkingDir creates a new temporary working directory
func NewWorkingDir() (*WorkDir, error) {
	twd := &WorkDir{}
	var err error

	twd.tmp, err = ioutil.TempDir("", "")
	if err != nil {
		return nil, errors.Wrap(err, "creating directory")
	}

	twd.cwd, err = os.Getwd()
	if err != nil {
		return nil, errors.Wrap(err, "getting current wd")
	}

	err = os.Chdir(twd.tmp)
	if err != nil {
		return nil, errors.Wrap(err, "changing cwd")
	}

	return twd, nil
}

// Cleanup restores the original working directory and cleans up the tmp dir
func (t *WorkDir) Cleanup() error {
	err := os.Chdir(t.cwd)
	if err != nil {
		return err
	}
	return os.RemoveAll(t.tmp)
}

// Join appends the relative path specified by file to the original working directory
func (t *WorkDir) Join(file string) string {
	return filepath.Join(t.cwd, file)
}
