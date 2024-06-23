package common

import (
	"os"
	"os/signal"
	"syscall"
)

var initFunctions = make([]func(), 0)
var exitFunctions = make([]func(), 0)

func InitCommon() {
	for _, apply := range initFunctions {
		apply()
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	go func(ch chan os.Signal) {
		<-ch
		for _, apply := range exitFunctions {
			apply()
		}
		os.Exit(0)
	}(ch)
}

func AddInitialized(apply func()) {
	initFunctions = append(initFunctions, apply)
}

func AddExited(apply func()) {
	exitFunctions = append(exitFunctions, apply)
}
