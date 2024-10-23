package main

import (
	"basil/libraries/humanize_filesize"
	"fmt"
	"math/rand"
)

func main() {
	v := rand.Int31n(1000000)
	fmt.Printf(`%d bytes = %s\n`, v, humanize_filesize.GetHumanizedFilesize(&v))
}
