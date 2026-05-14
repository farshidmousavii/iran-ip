package formatter

import (
	"bytes"
	"fmt"
)

type RawJSONFormatter struct{}

func (RawJSONFormatter) Name() string { return "rawjson" }

func (RawJSONFormatter) Format(v4, v6 []string, timestamp string) ([]File, error) {
	var buf bytes.Buffer

	all := append(v4, v6...)

	buf.WriteString("[\n")
	for i, s := range all {
		comma := ","
		if i == len(all)-1 {
			comma = ""
		}
		buf.WriteString(fmt.Sprintf("  \"%s\"%s\n", s, comma))
	}
	buf.WriteString("]\n")

	return []File{
		{Path: "raw/iran.json", Content: buf.Bytes()},
	}, nil
}
