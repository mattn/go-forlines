package forlines

import (
	"bufio"
	"bytes"
	"io"
)

func MustForLines(r io.Reader, f func(string) error) {
	err := ForLines(r, f)
	if err != nil {
		panic(err)
	}
}

func ForLines(r io.Reader, f func(string) error) error {
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
			err = f(buf.String())
			if err != nil {
				return err
			}
			buf.Reset()
		}
	}
	return nil
}
