package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func killServer(pidFile string) error {
	data, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return errors.Wrap(err, "can't open pid file (is server running?)")
	}

	errWhileRemovingFile := os.Remove(pidFile)

	if errWhileRemovingFile != nil {
		// We can go on if we fail here
		log.Printf("warning: can't remove pid file - %s", err)
	}

	strPID := strings.TrimSpace(string(data))
	pid, errConvertingDataToInt := strconv.Atoi(strPID)

	if errConvertingDataToInt != nil {
		return errors.Wrap(err, "bad process ID")
	}

	// Simulate kill
	fmt.Printf("killing server with pid=%d\n", pid)
	return nil
}

func main() {
	if err := killServer("server.pid"); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
