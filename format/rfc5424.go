package format

import (
	"bufio"

	"github.com/Xiol/syslogparser"
	"github.com/Xiol/syslogparser/rfc5424"
)

type RFC5424 struct{}

func (f *RFC5424) GetParser(line []byte) syslogparser.LogParser {
	return rfc5424.NewParser(line)
}

func (f *RFC5424) GetSplitFunc() bufio.SplitFunc {
	return nil
}
