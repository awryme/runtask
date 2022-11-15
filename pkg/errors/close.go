package errors

import "io"

func CloseInto(errp *error, closer io.Closer, msg string, args ...interface{}) {
	if errp == nil {
		return
	}
	if closer == nil {
		return
	}
	if *errp != nil {
		return
	}
	*errp = Wrap(closer.Close(), msg, args...)
}
