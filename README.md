# go-forlines

Another solution of https://github.com/Maki-Daisuke/go-lines

## Usage

```go
package main

import (
	"os"

	. "github.com/mattn/go-forlines"
)

func main() {
	err := ForLines(os.Stdin, func(line string) error {
		return do_something_with(line)
	})
	if err != nil {
		panic(err)
	}
}
```

```go
package main

import (
	"os"

	. "github.com/mattn/go-forlines"
)

func main() {
	MustForLines(os.Stdin, func(line string) error {
		return do_something_with(line)
	})
}
```

```go
package main

import (
	"os"

	. "github.com/mattn/go-forlines"
)

func main() {
	MustForLines(os.Stdin, do_something_with(line))
}
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a mattn)
