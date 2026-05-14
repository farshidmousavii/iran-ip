package formatter

import (
	"bytes"
	"fmt"
)

type SingboxFormatter struct{}

func (SingboxFormatter) Name() string { return "singbox" }

func (SingboxFormatter) Format(v4, v6 []string, timestamp string) ([]File, error) {
	var buf bytes.Buffer

	all := append(v4, v6...)

	buf.WriteString("{\n")
	buf.WriteString("  \"version\": 2,\n")
	buf.WriteString("  \"rules\": [\n")
	buf.WriteString("    {\n")
	buf.WriteString("      \"ip_cidr\": [\n")

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
		{Path: "sing-box/iran.json", Content: buf.Bytes()},
	}, nil
}
