package errors

import (
	"fmt"
	"io"
)

type wrapped struct {
	err error
	msg string
}

func (w *wrapped) Error() string { return w.msg + ": " + w.err.Error() }
func (w *wrapped) Cause() error  { return w.err }

// Unwrap provides compatibility for Go 1.13 error chains.
func (w *wrapped) Unwrap() error { return w.err }

func (w *wrapped) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v\n", w.Cause())
			io.WriteString(s, w.msg)
			return
		}
		fallthrough
	case 's', 'q':
		io.WriteString(s, w.Error())
	}
}

func Wrap(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	return &wrapped{err, msg}
}
