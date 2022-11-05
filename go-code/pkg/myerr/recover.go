package myerr

import (
	"runtime"

	"github.com/sirupsen/logrus"
)

// Recover recover info
func Recover(clean func()) {
	if err := recover(); err != nil {
		buf := make([]byte, 1024)
		for {
			n := runtime.Stack(buf, false)
			if n < len(buf) {
				buf = buf[:n]
				break
			}
			buf = make([]byte, 2*len(buf))
		}
		logrus.Errorf("%v\nstacktrace from panic: %s", err, string(buf))
		if clean != nil {
			clean()
		}
	}
}
