package formatter

import (
	"bytes"
	"fmt"
)

type ClashFormatter struct{}

func (ClashFormatter) Name() string { return "clash" }

func (ClashFormatter) Format(v4, v6 []string, timestamp string) ([]File, error) {
	var buf bytes.Buffer

	all := append(v4, v6...)

	buf.WriteString(fmt.Sprintf("# last fetch: %s\n", timestamp))
	buf.WriteString("payload:\n")
	for _, s := range all {
		buf.WriteString(fmt.Sprintf("  - '%s'\n", s))
	}

	return []File{
		{Path: "clash/iran.yaml", Content: buf.Bytes()},
	}, nil
}
