package formatter

import (
	"bytes"
	"fmt"
)

type XrayFormatter struct{}

func (XrayFormatter) Name() string { return "xray" }

func (XrayFormatter) Format(v4, v6 []string, timestamp string) ([]File, error) {
	var buf bytes.Buffer

	all := append(append([]string{}, v4...), v6...)

	buf.WriteString("{\n")
	buf.WriteString("  \"rules\": [\n")
	buf.WriteString("    {\n")
	buf.WriteString("      \"type\": \"field\",\n")
	buf.WriteString("      \"ip\": [\n")

	for i, s := range all {
		comma := ","
		if i == len(all)-1 {
			comma = ""
		}
		buf.WriteString(fmt.Sprintf("        \"%s\"%s\n", s, comma))
	}

	buf.WriteString("      ]\n")
	buf.WriteString("    }\n")
	buf.WriteString("  ]\n")
	buf.WriteString("}\n")

	return []File{
		{Path: "xray/iran.json", Content: buf.Bytes()},
	}, nil
}
