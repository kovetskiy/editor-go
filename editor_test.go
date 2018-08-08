package editor

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	test := assert.New(t)

	tempDir, err := ioutil.TempDir(os.TempDir(), "editor-go.test.")
	if err != nil {
		panic(err)
	}

	pathBinary := filepath.Join(tempDir, "editor")
	pathArgs := filepath.Join(tempDir, "args")
	pathContent := filepath.Join(tempDir, "content")

	expectedContent := time.Now().String()

	err = ioutil.WriteFile(pathBinary, []byte(`#!/bin/bash

echo "${@}" > `+pathArgs+`
echo -n "`+expectedContent+`" > $1
`), 0777)
	test.NoError(err)

	SetEditor(pathBinary)

	actualContents, err := Run(pathContent)
	test.NoError(err)
	test.Equal([]byte(expectedContent), actualContents)
}
