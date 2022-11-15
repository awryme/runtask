package stdio

import "os"

type Handles struct {
	Stdin, Stdout, Stderr *os.File
}

func DefaultHandles() Handles {
	return Handles{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
}
