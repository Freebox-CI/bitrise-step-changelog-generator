package main

import (
	"fmt"
	"os"
	"os/exec"
)

const DebugEnv = "debug"
const DebugKeyOk = "yes"

func main() {
	isDebug :=isDebug()
	prefixStrList := extractTypeList()
	entries := createEntries(prefixStrList)
	commitStrList := extractCommitList()
	fillCommitInfo(commitStrList, entries)
	unicodeResult := getBasicResult(entries)

	if isDebug {
		displayEntries(entries)
		fmt.Printf("%s", unicodeResult)
	}

	cmdLog, err := exec.Command("bitrise", "envman", "add", "--key", "CHANGELOG_BASIC", "--value", unicodeResult).CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to expose output with envman, error: %#v | output: %s", err, cmdLog)
		os.Exit(1)
	}else {
		os.Exit(0) //Step as "successful"
	}
}

func isDebug() bool {
	return os.Getenv(DebugEnv) == DebugKeyOk
}