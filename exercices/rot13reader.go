package main

import (
	"io"
	"os"
	"strings"
)

/*
 * Implements rot13Reader applying rot13 substitution
 *
 * rot13:
 * a-m -> n-z
 * n-z -> a-m
 */

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n,err := r.r.Read(b)

	for i,c := range b {
		b[i] = rot13(c)
	}
	return n, err
}

func rot13(b byte) byte {
	if (b >= 'A' && b <= 'M') || (b >= 'a' && b <= 'm') {
		return b+13
	}
	if (b >= 'N' && b <= 'Z') || (b >= 'n' && b <= 'z') {
		return b-13
	}
	return b
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}