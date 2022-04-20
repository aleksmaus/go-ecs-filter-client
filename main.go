package main

import (
	"fmt"
	"os"

	"github.com/elastic/go-ecs-filter/filter"
)

func exitIfError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
func main() {
	f, err := filter.New()
	exitIfError(err)
	defer f.Close()

	fmt.Println("sdsdf")
}
