package taskfile

import (
	"bufio"
	"fmt"
	"github.com/awryme/runtask/pkg/errors"
	"io"
	"strings"
)

func Decode(reader io.Reader) (*Taskfile, error) {
	data := &Taskfile{}

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}
		cmd, err := parseLine(line)
		if err != nil {
			return nil, errors.Wrap(err, "parse line")
		}
		data.Commands = append(data.Commands, cmd)
	}
	return data, errors.Wrap(scanner.Err(), "decode file")
}

func parseLine(line string) (*Command, error) {
	name, cmd, found := strings.Cut(line, ":")
	if !found {
		return nil, fmt.Errorf("line '%s' should have ':'", line)
	}
	name = strings.TrimSpace(name)
	cmd = strings.TrimSpace(cmd)

	return &Command{name, cmd}, nil
}
