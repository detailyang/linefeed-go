package linefeed

import "runtime"

var linefeed []byte

func init() {
	if runtime.GOOS == "windows" {
		linefeed = []byte("\r\n")
	} else if runtime.GOOS == "darwin" {
		linefeed = []byte("\r")
	} else {
		linefeed = []byte("\n")
	}
}

// Get gets the linefeed on the native platform
func Get() []byte {
	return linefeed
}
