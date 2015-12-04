package forlines

import (
	"bufio"
	"bytes"
	"io"
)

func MustForLines(r io.Reader, f func(string) bool) {
	err := ForLines(r, f)
	if err != nil {
		panic(err)
	}
}

func ForLines(r io.Reader, f func(string) bool) error {
	br := bufio.NewReader(r)
	var buf bytes.Buffer
	for {
		line, isPrefix, err := br.ReadLine()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		if _, err = buf.Write(line); err != nil {
			return err
		}
		if !isPrefix {
			if !f(buf.String()) {
				return nil
			}
			buf.Reset()
		}
	}
	return nil
}
