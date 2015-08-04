package main

import (
	"fmt"
	"os"

	"github.com/levenlabs/go-srvclient"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <hostname>\n", os.Args[0])
		os.Exit(1)
	}

	r, err := srvclient.SRV(os.Args[1])
	if err != nil {
		fmt.Printf("error resolving '%q': %s\n", os.Args[1], err)
		os.Exit(2)
	}

	fmt.Println(r)
}
