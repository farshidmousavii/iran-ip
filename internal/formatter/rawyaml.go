package formatter

import (
	"bytes"
	"fmt"
)

type RawYAMLFormatter struct{}

func (RawYAMLFormatter) Name() string { return "rawyaml" }

func (RawYAMLFormatter) Format(v4, v6 []string, timestamp string) ([]File, error) {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("# last fetch: %s\n", timestamp))

	all := append(append([]string{}, v4...), v6...)

	for _, s := range all {
		buf.WriteString(fmt.Sprintf("- \"%s\"\n", s))
	}

	return []File{
		{Path: "raw/iran.yaml", Content: buf.Bytes()},
	}, nil
}
