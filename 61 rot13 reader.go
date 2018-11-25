package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(p []byte) (n int, err error) {
	n, err = reader.r.Read(p)

	for i, val := range p {
		if val >= 'A' && val <= 'Z' {
			p[i] = (val-'A'+13)%26 + 'A'
		} else if val >= 'a' && val <= 'z' {
			p[i] = (val-'a'+13)%26 + 'a'
		}
	}
	return
}

func main() {
	s := strings.NewReader(
		"Lbh penpxrq gur pbqr!",
	)
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
