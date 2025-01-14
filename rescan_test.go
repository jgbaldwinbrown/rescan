package rescan

import (
	"testing"
	"regexp"
	"strings"
	"bufio"
	"fmt"
)

const in = `>one
gtcagtcagtcgtcagtca
tcagtca
cagacgtgtcagtcagt

cagtagtgtca
>two
agtcgtcagtcagtcaagtccagtgtcagtca
gact
cagtgtcagtcagtcagt
`

var faRe = regexp.MustCompile(`>[^\n]*\n[^>]*`)

func TestReSplitFunc(t *testing.T) {
	s := bufio.NewScanner(strings.NewReader(in))
	s.Split(ReSplitFunc(faRe))
	for i := 0; s.Scan(); i++ {
		if s.Err() != nil {
			t.Error(s.Err())
		}
		fmt.Printf("entry %v: %v\n", i, s.Text())
	}
}
