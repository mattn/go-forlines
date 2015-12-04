package forlines

import (
	"os"
	"reflect"
	"testing"
)

func TestLinesWithError(t *testing.T) {
	f, err := os.Open(`testdata/lines.txt`)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	expect := []string{
		"hoge",
		"fuga",
		"piyo",
		"",
		"foo bar baz",
	}
	got := []string{}

	err = ForLines(f, func(line string) error {
		got = append(got, line)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(expect, got) {
		t.Fatalf("expect %v but got %v", expect, got)
	}
}
