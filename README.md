# durafmt

durafmt is a tiny Go library that formats `time.Duration` strings into a human readable format.

```
go get github.com/hako/durafmt
```

# Why

If you've worked with `time.Duration` in Go, you most likely have come across this:

```
53m28.587093086s // :)
```

The above seems very easy to read, unless your duration looks like this:

```
354h22m3.24s // :S
```

# Usage

```go
package main

import (
	"fmt"	
	"github.com/hako/durafmt"
)

func main() {
	duration, err := durafmt.Parse("354h22m3.24s")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(duration) // 2 weeks 18 hours 22 minutes 3 seconds
	// duration.String() // String representation. "2 weeks 18 hours 22 minutes 3 seconds"

}
```

# Contributing

Contributions are welcome! Fork this repo and add your changes and submit a PR.

If you would like to fix a bug, add a feature or provide feedback you can do so in the issues section.

You can run tests by runnning `go test`. Running `go test; go vet; golint` is recommended.

durafmt is also tested against `gometalinter`.

# License

MIT