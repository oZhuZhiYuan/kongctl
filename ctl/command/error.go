package command

import (
	"fmt"
	"os"
)

const (
	// http://tldp.org/LDP/abs/html/exitcodes.html
	ExitSuccess = iota
	ExitError
	ExitBadConnection
	ExitInvalidInput // for txn, watch command
	ExitBadFeature   // provided a valid flag with an unsupported value
	ExitBadFlag
	ExitInterrupted
	ExitIO
	ExitBadArgs = 128
)

func ExitWithError(code int, err error) {
	fmt.Fprintln(os.Stderr, "Error:", err)
	os.Exit(code)
}
