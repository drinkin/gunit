package gunit

import (
	"fmt"
	"strconv"
	"strings"
)

func FormatFailureContext(lineNumber int, code string) string {
	lines := strings.Split(code, "\n")
	if len(lines) == 0 || lineNumber >= len(lines) {
		return ""
	}

	failure := []string{formatFailure(lines[lineNumber], lineNumber)}
	if len(lines) > lineNumber {
		failure = append(failure, format(lines[lineNumber+1], strconv.Itoa(lineNumber+1)))
	}

	for x := lineNumber - 1; x > 0; x-- {
		line := lines[x]
		if strings.HasPrefix(line, "func (") {
			failure = insert(failure, format(line, strconv.Itoa(x)))
			break
		} else /* if x == lineNumber-1 */ {
			failure = insert(failure, format(line, strconv.Itoa(x)))
			// } else if failure[0] != "Line .. \t| \t..." {
			// failure = insert(failure, "Line .. \t| \t...")
			// } else {
			// continue
		}
	}

	return "\n" + strings.Join(failure, "\n") + "\n\n"
}

func formatFailure(line string, number int) string {
	if strings.HasPrefix(line, "\t") {
		line = line[1:]
	}
	return fmt.Sprintf("%d: ----------> %s", number, line)
}
func format(line string, number string) string {
	return fmt.Sprintf("\t%s", line)
}

// From: https://github.com/golang/go/wiki/SliceTricks
func insert(lines []string, value string) []string {
	lines = append(lines, "")
	copy(lines[1:], lines[0:])
	lines[0] = value
	return lines
}
