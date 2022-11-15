package taskfile

import (
	"bufio"
	"fmt"
	"github.com/awryme/runtask/pkg/errors"
	"io"
)

func Encode(writer io.Writer, data *Taskfile) error {
	w := bufio.NewWriter(writer)
	for _, command := range data.Commands {
		_, err := w.WriteString(fmt.Sprintf("%s: %s\n", command.Name(), command.RawCommand()))
		if err != nil {
			return errors.Wrap(err, "write string to file")
		}
	}
	return w.Flush()
}
