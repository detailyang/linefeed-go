package linefeed

import "runtime"

var linefeed []byte

func init() {
	if runtime.GOOS == "windows" {
		linefeed = []byte("\r\n")
	} else if runtime.GOOS == "darwin" {
        // class macos use \r but the newer version is \n
        // See https://en.wikipedia.org/wiki/Newline
		linefeed = []byte("\n")
	} else {
		linefeed = []byte("\n")
	}
}

// Get gets the linefeed on the native platform
func Get() []byte {
	return linefeed
}
