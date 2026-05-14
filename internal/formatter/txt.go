package formatter

import (
	"bufio"
	"bytes"
	"fmt"
)

type TxtFormatter struct{}

func (TxtFormatter) Name() string { return "txt" }

func (TxtFormatter) Format(v4, v6 []string, timestamp string) ([]File, error) {
	var v4Buf, v6Buf bytes.Buffer

	if err := writeTXT(&v4Buf, v4, timestamp); err != nil {
		return nil, err
	}
	if err := writeTXT(&v6Buf, v6, timestamp); err != nil {
		return nil, err
	}

	return []File{
		{Path: "raw/ipv4.txt", Content: v4Buf.Bytes()},
		{Path: "raw/ipv6.txt", Content: v6Buf.Bytes()},
	}, nil
}

func writeTXT(buf *bytes.Buffer, subnets []string, timestamp string) error {
	writer := bufio.NewWriter(buf)
	fmt.Fprintln(writer, "# last fetch:", timestamp)
	for _, subnet := range subnets {
		fmt.Fprintln(writer, subnet)
	}
	return writer.Flush()
}
