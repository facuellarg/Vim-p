package utils

import (
	"os"
	"strconv"
)

func Debug() bool {
	debug, _ := strconv.ParseBool(os.Getenv("DEBUG"))
	return debug
}
