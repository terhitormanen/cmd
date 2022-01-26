package main_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/terhitormanen/cmd/model"
	main "github.com/terhitormanen/cmd/revel"
	"github.com/terhitormanen/cmd/utils"
)

// test the commands
func TestClean(t *testing.T) {
	a := assert.New(t)
	gopath := setup("revel-test-clean", a)

	t.Run("Clean", func(t *testing.T) {
		a := assert.New(t)
		c := newApp("clean-test", model.NEW, nil, a)
		main.Commands[model.NEW].RunWith(c)
		c.Index = model.TEST
		main.Commands[model.TEST].RunWith(c)
		a.True(utils.Exists(filepath.Join(gopath, "clean-test", "app", "tmp", "main.go")),
			"Missing main from path "+filepath.Join(gopath, "clean-test", "app", "tmp", "main.go"))
		c.Clean.ImportPath = c.ImportPath
		a.Nil(main.Commands[model.CLEAN].RunWith(c), "Failed to run clean-test")
		a.False(utils.Exists(filepath.Join(gopath, "clean-test", "app", "tmp", "main.go")),
			"Did not remove main from path "+filepath.Join(gopath, "clean-test", "app", "tmp", "main.go"))
	})
	if !t.Failed() {
		if err := os.RemoveAll(gopath); err != nil {
			a.Fail("Failed to remove test path")
		}
	}
}
