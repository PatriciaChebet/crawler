package main

import (
	"fmt"
	"os"
)

func main() {
	cliArguments := os.Args[1:]

	if len(cliArguments) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(cliArguments) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else {
		fmt.Printf("starting crawl of: %s", os.Args[1])
	}
}
