package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r.r.Read(b)
	if err != nil {
		return
	}
	for i, v := range(b) {
		switch {
		case ('a' <= v && v <= 'm') || ('A' <= v && v <= 'M'):
			b[i] = v + 13
		case ('n' <= v && v <= 'z') || ('N' <= v && v <= 'Z'):
			b[i] = v - 13
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
