package editor

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/reconquest/karma-go"
)

var (
	DefaultEditor = "vim"

	editor = ""
)

func SetEditor(name string) {
	editor = name
}

func getEditor() string {
	if editor != "" {
		return editor
	}

	env := os.Getenv("EDITOR")
	if env != "" {
		return env
	}

	return DefaultEditor
}

func Run(filename string) ([]byte, error) {
	cmd := exec.Command(getEditor(), filename)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		return nil, karma.Format(
			err,
			"unable to run editor: %s", getEditor(),
		)
	}

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, karma.Format(
			err,
			"unable to read contents of file",
		)
	}

	return contents, nil
}

func RunTemporary(dirPrefix, filename string) ([]byte, error) {
	dir, err := ioutil.TempDir(os.TempDir(), dirPrefix)
	if err != nil {
		return nil, karma.Format(
			err,
			"unable to create temporary directory",
		)
	}

	defer os.RemoveAll(dir)

	path := filepath.Join(dir, filename)

	return Run(path)
}
