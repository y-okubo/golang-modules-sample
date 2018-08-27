package main // import "github.com/y-okubo/golang-modules-sample"

import (
	"fmt"

	humanize "github.com/dustin/go-humanize"
)

func main() {
	fmt.Printf("That file is %s.\n", humanize.Bytes(82854982)) // That file is 83 MB.
}
