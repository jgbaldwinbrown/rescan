package rescan

import (
	"regexp"
	"bufio"
)

func ReSplitFunc(re *regexp.Regexp) func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		idxs := re.FindIndex(data)
		if (idxs == nil) {
			if atEOF {
				return 0, nil, bufio.ErrFinalToken
			}
			return 0, nil, nil
		}
		if idxs[1] == len(data) && !atEOF {
			return 0, nil, nil
		}
		if atEOF {
			return idxs[1], data[idxs[0]:idxs[1]], bufio.ErrFinalToken
		}
		return idxs[1], data[idxs[0]:idxs[1]], nil
	}
}
