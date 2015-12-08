package forlines

import (
	"bufio"
	"bytes"
	"io"
)

func Must(r io.Reader, f func(string) error) {
	err := Do(r, f)
	if err != nil {
		panic(err)
	}
}

func Do(r io.Reader, f func(string) error) error {
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
