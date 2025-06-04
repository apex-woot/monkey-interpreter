package dlog

import (
	"io"
	"log"
	"os"
)

var Debug = log.New(io.Discard, "[DEBUG] ", log.LstdFlags)

func Configure(showDebug bool) {
	if showDebug {
		Debug.SetOutput(os.Stdout)
	} else {
		// Silence all debug output.
		Debug.SetOutput(io.Discard)
	}
}
