package utils

import (
	"fmt"
	"os"
)

// HandleFailure wraps errors that might occur during the runs.
// It returns the error message (to STDERR) and exits.
func HandleFailure(err error) {
	if err != nil {
		errMessage := fmt.Sprintf("[ERROR] %v\n", err)
		fmt.Fprintf(os.Stderr, errMessage)
		os.Exit(1)
	}

}
