package utils

import (
	"regexp"
	"strings"

	"warehouse_rack/errors"
)

const (
	Tab          = "\t"
	Space        = " "
	NewLineDelim = "\n"
	EndLineDelim = '\n'
)

var regNoRegex = regexp.MustCompile(`^(([A-Za-z]){2}(|-)(?:[0-9]){1,2}(|-)(?:[A-Za-z]){1,2}(|-)([0-9]){1,4})$`)

func SplitCmdArguments(str string) (res []string, err error) {
	if strings.Contains(str, Tab) {
		err = errors.ErrInvalidTabSpace
		return
	}
	return strings.SplitN(str, Space, 2), nil
}
