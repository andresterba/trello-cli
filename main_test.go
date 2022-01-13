package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func buildTestExecutable() ([]byte, error) {
	command := exec.Command("go", "build", "-o", "test-trello-cli")
	output, err := command.CombinedOutput()

	return output, err
}

func removeExecutable() {
	os.Remove("gosdnc")
}

func checkCmdErr(test *testing.T, err error, output []byte) {
	if err != nil {
		fmt.Printf("Error: %s \n Output: %s", err, string(output))
		test.FailNow()
	}
}

func TestMain(t *testing.T) {
	defer removeExecutable()

	output, err := buildTestExecutable()
	checkCmdErr(t, err, output)

	runCommand := exec.Command("./test-trello-cli", "help")

	expectedOutput := `
trello-cli [command] [options]
commands:
    context - Set context for other commands.
    list - Print board and card ids.
    recurring - Print or add recurring todo's.
    version - Print current version.
    todo - Print current todo's.
    shopping-list - Interact with our shopping-list.
`

	output, err = runCommand.CombinedOutput()

	checkCmdErr(t, err, output)
	assert.Equal(t, expectedOutput, string(output), "they should be equal")
}
